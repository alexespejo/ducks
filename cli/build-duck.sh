go build -o duck.exe main.go
cp duck.exe duck.exe.bak
sudo mv duck.exe /usr/local/bin
chmod +x /usr/local/bin/duck 
mv duck.exe.bak duck.exe

touch ../src/run-astro.bash
cat > ../src/run-astro.bash << 'EOF'
#!/bin/bash

if [ $# -ne 1 ]; then
  echo "Error: This script requires exactly one argument"
  exit 1
fi

case "$1" in
  "dev")
    yarn astro dev
    ;;
  "build")
    yarn astro build
    ;;
  *)
    echo "Error: Invalid argument. Use 'dev', 'build', or 'preview'"
    exit 1
    ;;
esac
EOF

grep -q "^run-astro.bash$" ../.gitignore || echo "run-astro.bash" >> ../.gitignore

duck start