# ssh-add-keys

to check it, for example:
```shell
~$ ssh-keygen -t ed25519 -C "test key"  -P "Yungh%614Aeqt^%" -f test-keys/private1.key
~$ ssh-keygen -t ed25519 -C "test key"  -P "M^^^Aer**12<>" -f test-keys/private2.key
```
```shell
~$ make build
````
```shell
~$ bin/ssh-add-keys
````
```shell
~$ ssh-add -l
```
