## Special Thanks

Most of the locales have been taken and adapted from the Ruby Faker library by Benjamin Curties
https://github.com/stympy/faker under the MIT License (License added to this readme and folder).

## How locales work
Locales are just YAML files with a pseudo-template language.

Most of the fields are the regular arrays (using different formats for legibility). 
For example: 
`name: ["John", "Jack"]` 

There are, however, other fields a little bit more complex, like the following.
* Formatters. They format strings following a template
    * Digits `#`. Each `#` is a digit from 0-9. For example `##` could generate `34` and `(#-##)` could generate
    the string `(2-54)`
    * Field names. They are inside `{{}}`. They can refer to a field of the same YAML file
    and root. For example, if under `person` we see `{{first_name}}` the formatter will take a
    random value of the `first_name` key.
    It also can refer to another generator name. This means, if we are in `internet` generator
    but we want to use a `person` key, we can do it using it with a dot `{{person.first_name}}`. 
    Note that this behavior needs to be coded in the Generator if you plan to add a new reference to an
    external generator in the template.

You can modify and/or add new fields and new code for any generator!

## License

Copyright (c) 2007-2010 Benjamin Curtis

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.