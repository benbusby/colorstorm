fs = require 'fs'
module.exports =
  config:
    SelectSyntax:
      type: 'string',
      default: 'fire-spring-syntax',
      enum: ['fire-spring-syntax',
      'fire-spring-darker-syntax',
      'earthbound-syntax',
      'earthbound-darker-syntax',
      'dusty-dunes-syntax',
      'dusty-dunes-darker-syntax',
      'magicant-light-syntax',
      'cave-of-the-past-syntax',
      'zombie-threed-syntax',
      'zombie-threed-darker-syntax',
      'moonside-syntax'
      ]

atom.config.onDidChange 'earthbound-syntax.SelectSyntax', ({newValue, oldValue}) ->
    if (newValue)
        fs.createReadStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/themes/' + newValue + '/colors.less').pipe(fs.createWriteStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/styles/colors.less'));
    else
        fs.createReadStream(atom.packages.getPackageDirPaths() + 'earthbound-syntax/themes/fire-spring-syntax/colors.less').pipe(fs.createWriteStream(atom.packages.getPackageDirPaths() + '/earthbound-syntax/styles/colors.less'));
