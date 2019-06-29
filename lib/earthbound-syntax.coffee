fs = require 'fs'
module.exports =
  config:
    SelectSyntax:
      type: 'string',
      default: 'fire-spring-syntax',
      enum: ['fire-spring-syntax',
      'earthbound-syntax',
      'dusty-dunes-syntax',
      'magicant-light-syntax',
      'cave-of-the-past-light-syntax',
      'zombie-threed-syntax',
      'moonside-syntax'
      ]

atom.config.onDidChange 'earthbound-syntax.SelectSyntax', ({newValue, oldValue}) ->
    if (newValue)
        fs.createReadStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/atom/' + newValue + '/colors.less').pipe(fs.createWriteStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/styles/colors.less'));
    else
        fs.createReadStream(atom.packages.getPackageDirPaths() + 'earthbound-syntax/atom/fire-spring-syntax/colors.less').pipe(fs.createWriteStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/styles/colors.less'));
