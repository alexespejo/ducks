---
layout: "../../layouts/LayoutSingle.astro"
title: AA_Old_Code
---

```ts
export const saveSchedule = async (userID: string, rememberMe: boolean) => {
 logAnalytics({
  category: analyticsEnum.nav.title,
  action: analyticsEnum.nav.actions.SAVE_SCHEDULE,
  label: userID,
  value: rememberMe ? 1 : 0,
 });

 if (userID != null) {
  userID = userID.replace(/\s+/g, "");

  if (userID.length > 0) {
   if (rememberMe) {
    setLocalStorageUserId(userID);
   } else {
    removeLocalStorageUserId();
   }

   const scheduleSaveState = AppStore.schedule.getScheduleAsSaveState();

   if (
    isEmptySchedule(scheduleSaveState.schedules) &&
    !confirm(
     "You are attempting to save empty schedule(s). If this is unintentional, this may overwrite your existing schedules that haven't loaded yet!"
    )
   ) {
    return;
   }

   try {
    await trpc.users.saveUserData.mutate({
     id: userID,
     data: {
      id: userID,
      userData: scheduleSaveState,
     },
    });

    openSnackbar(
     "success",
     `Schedule saved under username "${userID}". Don't forget to sign up for classes on WebReg!`
    );
    AppStore.saveSchedule();
   } catch (e) {
    if (e instanceof TRPCError) {
     openSnackbar(
      "error",
      `Schedule could not be saved under username "${userID}`
     );
    } else {
     openSnackbar("error", "Network error or server is down.");
    }
   }
  }
 }
};
```

```ts
export const loadSchedule = async (userId: string, rememberMe: boolean) => {
 logAnalytics({
  category: analyticsEnum.nav.title,
  action: analyticsEnum.nav.actions.LOAD_SCHEDULE,
  label: userId,
  value: rememberMe ? 1 : 0,
 });
 if (
  userId != null &&
  (!AppStore.hasUnsavedChanges() ||
   window.confirm(
    `Are you sure you want to load a different schedule? You have unsaved changes!`
   ))
 ) {
  userId = userId.replace(/\s+/g, "");
  if (userId.length > 0) {
   if (rememberMe) {
    setLocalStorageUserId(userId);
   } else {
    removeLocalStorageUserId();
   }

   try {
    const res = await trpc.users.getUserData.query({ userId });
    const scheduleSaveState = res && "userData" in res ? res.userData : res;

    if (scheduleSaveState == null) {
     openSnackbar("error", `Couldn't find schedules for username "${userId}".`);
    } else if (await AppStore.loadSchedule(scheduleSaveState)) {
     openSnackbar("success", `Schedule for username "${userId}" loaded.`);
    } else {
     AppStore.loadSkeletonSchedule(scheduleSaveState);
     openSnackbar(
      "error",
      `Network error loading course information for "${userId}". 	              
                        If this continues to happen, please submit a feedback form.`
     );
    }
   } catch (e) {
    console.error(e);
    openSnackbar(
     "error",
     `Failed to load schedules. If this continues to happen, please submit a feedback form.`
    );
   }
  }
 }
};
```

```ts
import {
 Button,
 Dialog,
 DialogActions,
 DialogContent,
 DialogContentText,
 DialogTitle,
 TextField,
 CircularProgress,
 Checkbox,
 FormControlLabel,
} from "@material-ui/core";
import { CloudDownload, Save } from "@material-ui/icons";
import { LoadingButton } from "@mui/lab";
import { ChangeEvent, PureComponent, useEffect, useState } from "react";

import actionTypesStore from "$actions/ActionTypesStore";
import { loadSchedule, saveSchedule } from "$actions/AppStoreActions";
import { getLocalStorageUserId } from "$lib/localStorage";
import AppStore from "$stores/AppStore";
import { useThemeStore } from "$stores/SettingsStore";

interface LoadSaveButtonBaseProps {
 action: typeof saveSchedule;
 actionName: "Save" | "Load";
 disabled: boolean;
 loading: boolean;
 colorType: "primary" | "secondary";
 id?: string;
}

interface LoadSaveButtonBaseState {
 isOpen: boolean;
 userID: string;
 rememberMe: boolean;
}

interface SaveLoadIconProps {
 loading: boolean;
 actionName: "Save" | "Load";
}

function SaveLoadIcon(props: SaveLoadIconProps) {
 return props.loading ? (
  <CircularProgress size={20} color="inherit" />
 ) : props.actionName === "Save" ? (
  <Save />
 ) : (
  <CloudDownload />
 );
}

class LoadSaveButtonBase extends PureComponent<
 LoadSaveButtonBaseProps,
 LoadSaveButtonBaseState
> {
 state: LoadSaveButtonBaseState = {
  isOpen: false,
  userID: "",
  rememberMe: true,
 };

 handleOpen = () => {
  this.setState({ isOpen: true });
  if (typeof Storage !== "undefined") {
   const userID = getLocalStorageUserId();
   if (userID !== null) {
    this.setState({ userID: userID });
   }
  }
 };

 handleClose = (wasCancelled: boolean) => {
  if (wasCancelled)
   this.setState({ isOpen: false }, () => {
    document.removeEventListener("keydown", this.enterEvent, false);
    this.setState({ userID: "" });
   });
  else
   this.setState({ isOpen: false }, () => {
    document.removeEventListener("keydown", this.enterEvent, false);
    // this `void` is for eslint "no floating promises"
    void this.props.action(this.state.userID, this.state.rememberMe);
    this.setState({ userID: "" });
   });
 };

 handleToggleRememberMe = (event: ChangeEvent<HTMLInputElement>) => {
  this.setState({ rememberMe: event.target.checked });
 };

 componentDidUpdate(_prevProps: unknown, prevState: LoadSaveButtonBaseState) {
  if (!prevState.isOpen && this.state.isOpen)
   document.addEventListener("keydown", this.enterEvent, false);
  else if (prevState.isOpen && !this.state.isOpen)
   document.removeEventListener("keydown", this.enterEvent, false);
 }

 enterEvent = (event: KeyboardEvent) => {
  const charCode = event.which ? event.which : event.keyCode;

  if (charCode === 13 || charCode === 10) {
   event.preventDefault();
   this.handleClose(false);

   return false;
  }
 };

 render() {
  return (
   <>
    <LoadingButton
     id={this.props.id}
     onClick={this.handleOpen}
     color="inherit"
     startIcon={
      <SaveLoadIcon
       loading={this.props.loading}
       actionName={this.props.actionName}
      />
     }
     disabled={this.props.disabled}
     loading={false}
    >
     {this.props.actionName}
    </LoadingButton>
    <Dialog open={this.state.isOpen} onClose={this.handleClose}>
     <DialogTitle>{this.props.actionName}</DialogTitle>
     <DialogContent>
      <DialogContentText>
       Enter your unique user ID here to {this.props.actionName.toLowerCase()}{" "}
       your schedule.
      </DialogContentText>
      <DialogContentText style={{ color: "red" }}>
       Make sure the user ID is unique and secret, or someone else can overwrite
       your schedule.
      </DialogContentText>
      <TextField
       // eslint-disable-next-line jsx-a11y/no-autofocus
       autoFocus
       margin="dense"
       label="Unique User ID"
       type="text"
       fullWidth
       placeholder="Enter here"
       value={this.state.userID}
       onChange={(event) => this.setState({ userID: event.target.value })}
      />
      <FormControlLabel
       control={
        <Checkbox
         checked={this.state.rememberMe}
         onChange={this.handleToggleRememberMe}
         color="primary"
        />
       }
       label="Remember Me (Uncheck on shared computers)"
      />
     </DialogContent>
     <DialogActions>
      <Button
       onClick={() => this.handleClose(true)}
       color={this.props.colorType}
      >
       {"Cancel"}
      </Button>
      <Button
       onClick={() => this.handleClose(false)}
       color={this.props.colorType}
      >
       {this.props.actionName}
      </Button>
     </DialogActions>
    </Dialog>
   </>
  );
 }
}

const LoadSaveScheduleFunctionality = () => {
 const isDark = useThemeStore((store) => store.isDark);

 const [loading, setLoading] = useState(false);
 const [saving, setSaving] = useState(false);
 const [skeletonMode, setSkeletonMode] = useState(AppStore.getSkeletonMode());

 const loadScheduleAndSetLoading = async (
  userID: string,
  rememberMe: boolean
 ) => {
  setLoading(true);
  await loadSchedule(userID, rememberMe);
  setLoading(false);
 };

 const saveScheduleAndSetLoading = async (
  userID: string,
  rememberMe: boolean
 ) => {
  setSaving(true);
  await saveSchedule(userID, rememberMe);
  setSaving(false);
 };

 useEffect(() => {
  const handleSkeletonModeChange = () => {
   setSkeletonMode(AppStore.getSkeletonMode());
  };

  AppStore.on("skeletonModeChange", handleSkeletonModeChange);

  return () => {
   AppStore.off("skeletonModeChange", handleSkeletonModeChange);
  };
 }, []);

 useEffect(() => {
  if (typeof Storage !== "undefined") {
   const savedUserID = getLocalStorageUserId();

   if (savedUserID != null) {
    // this `void` is for eslint "no floating promises"
    void loadScheduleAndSetLoading(savedUserID, true);
   }
  }
 }, []);

 useEffect(() => {
  const handleAutoSaveStart = () => setSaving(true);
  const handleAutoSaveEnd = () => setSaving(false);

  actionTypesStore.on("autoSaveStart", handleAutoSaveStart);
  actionTypesStore.on("autoSaveEnd", handleAutoSaveEnd);

  return () => {
   actionTypesStore.off("autoSaveStart", handleAutoSaveStart);
   actionTypesStore.off("autoSaveEnd", handleAutoSaveEnd);
  };
 }, []);

 return (
  <div
   id="load-save-container"
   style={{ display: "flex", flexDirection: "row" }}
  >
   <LoadSaveButtonBase
    id="save-button"
    actionName={"Save"}
    action={saveScheduleAndSetLoading}
    disabled={loading}
    loading={saving}
    colorType={isDark ? "secondary" : "primary"}
   />
   <LoadSaveButtonBase
    id="load-button"
    actionName={"Load"}
    action={loadScheduleAndSetLoading}
    disabled={skeletonMode}
    loading={loading}
    colorType={isDark ? "secondary" : "primary"}
   />
  </div>
 );
};

export default LoadSaveScheduleFunctionality;
```

```ts
    static async upsertGuestUserData(db: DatabaseOrTransaction, userData: User): Promise<string> {
        return db.transaction(async (tx) => {
            const userId = await this.createGuestUserOptional(tx, userData.id);

            if (userId === undefined) {
                throw new Error(`Failed to create guest user for ${userData.id}`);
            }

            // Add schedules and courses
            const scheduleIds = await this.upsertSchedulesAndContents(tx, userId, userData.userData.schedules);

            // Update user's current schedule index
            const scheduleIndex = userData.userData.scheduleIndex;

            const currentScheduleId =
                scheduleIndex === undefined || scheduleIndex >= scheduleIds.length ? null : scheduleIds[scheduleIndex];

            if (currentScheduleId !== null) {
                await tx.update(users).set({ currentScheduleId: currentScheduleId }).where(eq(users.id, userId));
            }

            return userId;
        });
    }
```
