## Creating / Modifying Theme Templates

Reference the following table when creating or modifying theme templates:

| Value                   | Description                                           | Auto Generated? |
|-------------------------|-------------------------------------------------------|-----------------|
| `{{ .Background }}`     | The theme's primary background color                  |                 |
| `{{ .BackgroundAlt1 }}` | A 5% lighter\* version of `{{ .Background }}`         | ✅               |
| `{{ .BackgroundAlt2 }}` | A 10% lighter\* version of `{{ .Background }}`        | ✅               |
| `{{ .Foreground }}`     | The theme's primary foreground color                  |                 |
| `{{ .ForegroundAlt }}`  | A 10% darker\* version of `{{ .Foreground }}`         | ✅               |
| `{{ .Function }}`       | The color used for functions and methods              |                 |
| `{{ .Constant }}`       | The color used for constant values                    |                 |
| `{{ .Keyword }}`        | The color used for keywords (`if`, `struct`, etc)     |                 |
| `{{ .Comment }}`        | The color used for comments                           |                 |
| `{{ .Number }}`         | The color used for numeric values                     |                 |
| `{{ .String }}`         | The color used for string values                      |                 |
| `{{ .Type }}`           | The color used for data types (`int`, `float64`, etc) |                 |
| `{{ .Name }}`           | The theme name (may include spaces, punctuation, etc) |                 |
| `{{ .ID }}`             | A sanitized version of the theme name                 |                 |
| `{{ .Author }}`         | The user provided author name (can be empty)          |                 |
| `{{ .DarkOrLight }}`    | `light` for light themes, `dark` for dark themes      |                 |

<sup>* - This assumes the theme is a dark syntax theme. For light syntax themes, the X% lighter colors
would be darker, and the X% darker colors would be lighter.</sup>