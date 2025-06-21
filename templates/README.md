## Creating / Modifying Theme Templates

Reference the following table when creating or modifying theme templates:

| Value                   | Description                                               | Auto Generated? |
|-------------------------|-----------------------------------------------------------|-----------------|
| `{{ .Background }}`     | The theme's primary background hex color                  |                 |
| `{{ .BackgroundAlt1 }}` | A 5% lighter\* version of `{{ .Background }}`             | ✅               |
| `{{ .BackgroundAlt2 }}` | A 10% lighter\* version of `{{ .Background }}`            | ✅               |
| `{{ .Foreground }}`     | The theme's primary foreground hex color                  |                 |
| `{{ .ForegroundAlt }}`  | A 10% darker\* version of `{{ .Foreground }}`             | ✅               |
| `{{ .Function }}`       | The hex color used for functions and methods              |                 |
| `{{ .Constant }}`       | The hex color used for constant values                    |                 |
| `{{ .Keyword }}`        | The hex color used for keywords (`if`, `struct`, etc)     |                 |
| `{{ .Comment }}`        | The hex color used for comments                           |                 |
| `{{ .Number }}`         | The hex color used for numeric values                     |                 |
| `{{ .String }}`         | The hex color used for string values                      |                 |
| `{{ .Type }}`           | The hex color used for data types (`int`, `float64`, etc) |                 |
| `{{ .Name }}`           | The theme name (may include spaces, punctuation, etc)     |                 |
| `{{ .ID }}`             | A sanitized version of the theme name                     |                 |
| `{{ .Author }}`         | The user provided author name (can be empty)              |                 |
| `{{ .DarkOrLight }}`    | `light` for light themes, `dark` for dark themes          |                 |

<sup>* - This assumes the theme is a dark syntax theme. For light syntax themes, the X% lighter colors
would be darker, and the X% darker colors would be lighter.</sup>

**Note:** All hex color values can be suffixed with `X256` to get the XTerm 256 color value 
(for example: `{{ .BackgroundX256}}` or `{{ .FunctionX256 }}`).