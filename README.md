# SSH-Scanner
Simple SSH vulnerability scanner based on SSH Harvester.

## TODO

1. Refactor ssh-harvester code.
2. Grab all host keys. It depends on the supported ciphersuites. For example OpenSSH usually has four different keys (RSA, ECDAS, ED25... etc.).
3. Read defaults from the config file.
