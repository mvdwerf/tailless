# Tailless

This GO package is, as the name suggests, a combination of TailWind CSS and the LESS preprocessor.

## Usage

```go
func main() {
    err := tailless.Parse("style.less", "style.css")
    if err != nil {
        fmt.Println(err)
    }
}
```

## Example less file

```less
div
{
    .bg-neutral-100;
    .text-emerald-800;
}

a
{
    .text-white;
    .bg-emerald-700;

    &:hover
    {
        .bg-emerald-600;
    }
}

```

## Supported less features

```less
// Variables
@primary: #123456;

div
{
    background-color: @primary;
}

// Nesting
div
{
    &.red
    {
        background-color: red;
    }
}

// Mixins
.my_mixin
{
    color: red;
    background-color: yellow;
}

div
{
    .my_mixin;
}

// Tailwind properties are used like mixins:
div
{
    .bg-neutral-100;
    .text-cyan-800;
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
