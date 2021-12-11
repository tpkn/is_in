# Is in
A CLI tool for checking whether the parameters from first string are present in second string

## Usage

```text
is_in [-options] <string1> <string2>
```

## Options

```text
  -i, --ignorecase    Ignore case (default: false)
  -d, --delimiter     Parameters delimiter (default: ",")
  -p, --prepare       Prepare parameters before comparison (default: false): trim spaces and replace wrapping single/double quotes
  -w, --why           Print comparsison details (default: false)
  --help              Help
  --version           Print current version
```


## Shell usage
```shell
is_in ",f,g,b , ,c" "a,b,c"
# => "false"

is_in -why ",f,g,b , ,c" "a,b,c"
# =>
# Failed because:
# 1. 'f' is missing
# 2. 'g' is missing
# 3. 'b ' is missing
```

```shell
ok=$(is_in -p "a, b, c" "d,c")
if [ "$ok" = "false" ]; then
   echo "Error!"
   exit
fi
```

