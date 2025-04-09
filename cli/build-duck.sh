go build -o duck main.go
cp duck duck.bak
sudo mv duck /usr/local/bin
chmod +x /usr/local/bin/duck 
mv duck.bak duck.exe

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