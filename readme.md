# Box Drawing

## Companents
### Title
```go
title := NewTitle("test")
fmt.Println(title.Draw())
```
returns ```Test```

### Text
```go
text := NewText("test")
fmt.Println(text.Draw())
```
returns ```test```

### Box
```go
title := NewTitle("test")
bd := box.NewBox(title, box.Lines{'═', '╗', '║', '╝', '═', '╚', '║', '╔'})
fmt.Println(bd)
```
returns
```
╔════╗
║Test║
╚════╝
```

### Horizontal sequence
```go
	s := NewSequence(
		NewText("cell 1"),
		NewText("cell 2"),
		NewText("cell 3"),
		SequenceLines{'═', '╗', '║', '╝', '═', '╚', '║', '╔', '╦', '║', '╩'},
	)
	fmt.Println(s)
```
returns
```
╔══════╦══════╦══════╗
║cell 1║cell 2║cell 3║
╚══════╩══════╩══════╝
```

## Todo
- Table
- Load from SCV
- Import to HTML