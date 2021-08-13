# Is in
A CLI tool for checking whether the parameters from first string are present in second string

## Usage

```text
is_in [-options] <string1> <string2>
```

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



## Options

```text
-d (string)
      Parameter delimiter (default: ",")
-i	
      Ignore case (default: false)
-p	
      Prepare parameters before comparison (trim spaces and replace wrapping single/double quotes) (default: false)
-why
      Print comparison details
-v
      Print current version number
```


