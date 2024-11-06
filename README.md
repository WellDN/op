# Compile Bootloader


```

nasm -f bin boot.asm -o boot.bin

```
# Compile Kernel

```

GOOS=linux GOARCH=386 go build -o kernel.bin kernel.go

```

# Link Bootloader and Kernel

```

cat boot.bin kernel.bin > os-image.bin

```

# Create a Bootable ISO with GRUB

1. Create a directory structure for GRUB:

```bash

mkdir -p iso/boot/grub
cp os-image.bin iso/boot/

```

2. Create a grub.cfg file in iso/boot/grub with the following content:

```

set timeout=0
set default=0

menuentry "My Go OS" {
    multiboot /boot/os-image.bin
    boot
}

```

3. Build ISO image:

```

grub-mkrescue -o my-go-os.iso iso/


```

# Test OS with QEMU

```

qemu-system-i386 -cdrom my-go-os.iso -boot d


```


