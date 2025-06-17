package templates

import (
	"fmt"
	"text/template"
)

func GetVSCodeThemeFileName(name string) string {
	return fmt.Sprintf("%s.json", name)
}

func GetVSCodeThemeTemplate() (*template.Template, error) {
	return template.New("vscode_template").Parse(VSCodeTemplate)
}

const VSCodeTemplate = `{
  "name": "{{ .Name }}",
  "tokenColors": [
    {
      "name": "Global settings",
      "settings": {
        "background": "{{ .Background }}",
        "foreground": "{{ .Foreground }}"
      }
    },
    {
      "name": "String",
      "scope": "string",
      "settings": {
        "foreground": "{{ .String }}"
      }
    },
    {
      "name": "String Escape",
      "scope": "constant.character.escape, text.html constant.character.entity.named, punctuation.definition.entity.html",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Boolean",
      "scope": "constant.language.boolean",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Number",
      "scope": "constant.numeric",
      "settings": {
        "foreground": "{{ .Number }}"
      }
    },
    {
      "name": "Variable Instances",
      "scope": [
        "variable.instance",
        "variable.other.instance",
        "variable.readwrite.instance",
        "variable.other.readwrite.instance",
        "variable.other.property",
        "meta.definition.variable",
        "support.constant",
        "support.variable"
      ],
      "settings": {
        "foreground": "{{ .Foreground }}"
      }
    },
    {
      "name": "Identifier",
      "scope": "variable, support.class, entity.name.function",
      "settings": {
        "foreground": "{{ .Foreground }}"
      }
    },
    {
      "name": "Keyword",
      "scope": "keyword, modifier, variable.language.this, support.type.object, constant.language",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Function call",
      "scope": "entity.name.function, support.function",
      "settings": {
        "foreground": "{{ .Function }}"
      }
    },
    {
      "name": "Storage",
      "scope": "storage.type, storage.modifier",
      "settings": {
        "foreground": "{{ .Keyword }}"
      }
    },
    {
      "name": "Modules",
      "scope": "support.module, support.node",
      "settings": {
        "foreground": "{{ .Constant }}",
        "fontStyle": "italic"
      }
    },
    {
      "name": "Type",
      "scope": "support.type",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "Type",
      "scope": "entity.name.type, entity.other.inherited-class",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "Comment",
      "scope": "comment",
      "settings": {
        "foreground": "{{ .Comment }}",
        "fontStyle": "italic"
      }
    },
    {
      "name": "Class",
      "scope": "entity.name.type.class",
      "settings": {
        "foreground": "{{ .Type }}",
        "fontStyle": "underline"
      }
    },
    {
      "name": "Class variable",
      "scope": "variable.object.property, meta.field.declaration entity.name.function",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "Class method",
      "scope": "meta.definition.method entity.name.function",
      "settings": {
        "foreground": "{{ .Function }}"
      }
    },
    {
      "name": "Function definition",
      "scope": "meta.function entity.name.function",
      "settings": {
        "foreground": "{{ .Function }}"
      }
    },
    {
      "name": "Template expression",
      "scope": "template.expression.begin, template.expression.end, punctuation.definition.template-expression.begin, punctuation.definition.template-expression.end",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Reset embedded/template expression colors",
      "scope": "meta.embedded, source.groovy.embedded, meta.template.expression",
      "settings": {
        "foreground": "{{ .Foreground }}"
      }
    },
    {
      "name": "YAML key",
      "scope": "entity.name.tag.yaml",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "JSON key",
      "scope": "meta.object-literal.key, meta.object-literal.key string, support.type.property-name.json",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "JSON constant",
      "scope": "constant.language.json",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "CSS class",
      "scope": "entity.other.attribute-name.class",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "CSS ID",
      "scope": "entity.other.attribute-name.id",
      "settings": {
        "foreground": "{{ .String }}"
      }
    },
    {
      "name": "CSS tag",
      "scope": "source.css entity.name.tag",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "HTML tag outer",
      "scope": "meta.tag, punctuation.definition.tag",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "HTML tag inner",
      "scope": "entity.name.tag",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "HTML tag attribute",
      "scope": "entity.other.attribute-name",
      "settings": {
        "foreground": "{{ .Function }}"
      }
    },
    {
      "name": "Markdown heading",
      "scope": "markup.heading",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Markdown link text",
      "scope": "text.html.markdown meta.link.inline, meta.link.reference",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Markdown list item",
      "scope": "text.html.markdown beginning.punctuation.definition.list",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "Markdown italic",
      "scope": "markup.italic",
      "settings": {
        "foreground": "{{ .Constant }}",
        "fontStyle": "italic"
      }
    },
    {
      "name": "Markdown bold",
      "scope": "markup.bold",
      "settings": {
        "foreground": "{{ .Constant }}",
        "fontStyle": "bold"
      }
    },
    {
      "name": "Markdown bold italic",
      "scope": "markup.bold markup.italic, markup.italic markup.bold",
      "settings": {
        "foreground": "{{ .Constant }}",
        "fontStyle": "italic bold"
      }
    },
    {
      "name": "Markdown code block",
      "scope": "markup.fenced_code.block.markdown punctuation.definition.markdown",
      "settings": {
        "foreground": "{{ .String }}"
      }
    },
    {
      "name": "Markdown inline code",
      "scope": "markup.inline.raw.string.markdown",
      "settings": {
        "foreground": "{{ .String }}"
      }
    },
    {
      "name": "INI property name",
      "scope": "keyword.other.definition.ini",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "INI section title",
      "scope": "entity.name.section.group-title.ini",
      "settings": {
        "foreground": "{{ .Constant }}"
      }
    },
    {
      "name": "C# class",
      "scope": "source.cs meta.class.identifier storage.type",
      "settings": {
        "foreground": "{{ .Type }}",
        "fontStyle": "underline"
      }
    },
    {
      "name": "C# class method",
      "scope": "source.cs meta.method.identifier entity.name.function",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "C# function call",
      "scope": "source.cs meta.method-call meta.method, source.cs entity.name.function",
      "settings": {
        "foreground": "{{ .Function }}"
      }
    },
    {
      "name": "C# type",
      "scope": "source.cs storage.type",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "C# return type",
      "scope": "source.cs meta.method.return-type",
      "settings": {
        "foreground": "{{ .Type }}"
      }
    },
    {
      "name": "C# preprocessor",
      "scope": "source.cs meta.preprocessor",
      "settings": {
        "foreground": "{{ .Keyword }}"
      }
    },
    {
      "name": "C# namespace",
      "scope": "source.cs entity.name.type.namespace",
      "settings": {
        "foreground": "{{ .Foreground }}"
      }
    },
    {
      "name": "Global settings",
      "settings": {
        "background": "{{ .Background }}",
        "foreground": "{{ .Foreground }}"
      }
    }
  ],
  "colors": {
    "focusBorder": "{{ .Constant }}",
    "foreground": "{{ .Foreground }}",
    "button.background": "{{ .Constant }}",
    "button.foreground": "#000000",
    "dropdown.background": "{{ .BackgroundAlt2 }}",
    "input.background": "{{ .BackgroundAlt2 }}",
    "inputOption.activeBorder": "{{ .Constant }}",
    "list.activeSelectionBackground": "{{ .BackgroundAlt2 }}80",
    "list.activeSelectionForeground": "#FFFFFF",
    "list.dropBackground": "{{ .Constant }}80",
    "list.focusBackground": "{{ .Constant }}80",
    "list.focusForeground": "#FFFFFF",
    "list.highlightForeground": "{{ .Constant }}",
    "list.hoverBackground": "#FFFFFF1a",
    "list.inactiveSelectionBackground": "#FFFFFF33",
    "activityBar.background": "{{ .BackgroundAlt2 }}",
    "activityBar.foreground": "{{ .Foreground }}",
    "activityBar.dropBackground": "{{ .BackgroundAlt2 }}80",
    "activityBarBadge.background": "{{ .BackgroundAlt2 }}",
    "activityBarBadge.foreground": "{{ .Keyword }}",
    "badge.background": "{{ .BackgroundAlt2 }}",
    "badge.foreground": "{{ .Keyword }}",
    "sideBar.background": "{{ .BackgroundAlt1 }}",
    "sideBar.foreground": "{{ .Foreground }}",
    "sideBarSectionHeader.background": "{{ .BackgroundAlt2 }}",
    "editorGroup.dropBackground": "{{ .Constant }}80",
    "editorGroup.focusedEmptyBorder": "{{ .Constant }}",
    "editorGroupHeader.tabsBackground": "{{ .Background }}",
    "tab.border": "#00000033",
    "tab.activeBorder": "{{ .Constant }}",
    "tab.inactiveBackground": "{{ .BackgroundAlt1 }}",
    "tab.activeModifiedBorder": "{{ .Constant }}",
    "tab.inactiveModifiedBorder": "#908900",
    "tab.unfocusedActiveModifiedBorder": "#c0b700",
    "tab.unfocusedInactiveModifiedBorder": "#908900",
    "editor.background": "{{ .Background }}",
    "editor.foreground": "{{ .Foreground }}",
    "editor.selectionBackground": "{{ .BackgroundAlt2 }}",
    "editor.selectionHighlightBackground": "{{ .BackgroundAlt2 }}",
    "editorLineNumber.foreground": "{{ .ForegroundAlt }}dd",
    "editorLineNumber.activeForeground": "{{ .Keyword }}ff",
    "editor.lineHighlightBorder": "#FFFFFF1a",
    "editor.rangeHighlightBackground": "#FFFFFF0d",
    "editorWidget.background": "{{ .Background }}",
    "editorHoverWidget.background": "{{ .Background }}",
    "editorMarkerNavigation.background": "{{ .Background }}",
    "peekView.border": "{{ .Constant }}",
    "peekViewEditor.background": "{{ .BackgroundAlt1 }}",
    "peekViewResult.background": "{{ .Background }}",
    "peekViewTitle.background": "{{ .Background }}",
    "panel.background": "{{ .Background }}",
    "panel.border": "#FFFFFF1a",
    "panelTitle.activeBorder": "{{ .Foreground }}80",
    "panelTitle.inactiveForeground": "{{ .Foreground }}80",
    "statusBar.background": "{{ .BackgroundAlt1 }}",
    "statusBar.foreground": "{{ .Foreground }}",
    "statusBar.debuggingBackground": "{{ .Constant }}",
    "statusBar.debuggingForeground": "#000000",
    "statusBar.noFolderBackground": "{{ .BackgroundAlt1 }}",
    "statusBarItem.activeBackground": "{{ .Constant }}80",
    "statusBarItem.hoverBackground": "#FFFFFF1a",
    "statusBarItem.remoteBackground": "{{ .Constant }}",
    "statusBarItem.remoteForeground": "#000000",
    "titleBar.activeBackground": "{{ .BackgroundAlt1 }}",
    "pickerGroup.border": "#FFFFFF1a",
    "debugToolBar.background": "{{ .BackgroundAlt1 }}",
    "selection.background": "{{ .BackgroundAlt2 }}"
  }
}`
