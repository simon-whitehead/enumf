## enumf

Package enumf is a simple wrapper around an "enumeration" flag type. It attempts to make it cleaner to combine multiple flags together and check for individual flags that are set within it.

#### Example

A simple example would be:

```
import "github.com/simon-whitehead/enumf"

const (
	None enumf.E = 1 << iota
	Read
	Write
	Admin = 32768
)

func main() {
	permissions := enumf.E(Read | Write)

	if permissions.HasFlag(Write) {
		fmt.Println("Write permission available")
	}

	if permissions.HasFlag(Read) {
		fmt.Println("Read permission available")
	}

	if permissions.HasFlag(Admin) {
		fmt.Println("Admin powers")
	}
} 
```

Where the output is:

```
Write permission available
Read permission available
```

Setting flags can be accomplished like this:
```
permissions.SetFlag(Admin)

if permissions.HasFlag(Admin) {
    fmt.Println("Admin powers")
}
```
Where the output would be:
```
Admin powers
```

### Contributing
I honestly can't think of what else to add to this tiny tiny package. It just helps me clean up bit flag processing and I want to keep it narrowly focused to this task. If, however, you feel it could benefit from something else, feel free to submit a PR.

### License

Do what you please with it :)
