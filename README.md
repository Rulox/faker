
<p align="center"><img src="doc/faker_logo.png" width="250"></p>

# Faker (WIP, do not use this library in production)
Faker is a Golang library that generates all type of fake data. Including localized data.

This library is inspired by same libs in other languages like Perl's, ruby's and PHP's Faker.

Faker has been built and tested with Go >= 1.10

### Installing
Just use go get

`go get -u github.com/rulox/faker`

Or add the library to your project and use `dep ensure`

### Usage
The main struct `generator/Faker` provides all the usability for you in order to create the fake data.
By default, the data generated has no locale/language set (although some basic English data would be return).

`Faker` would have different number of generators (for Misc data, Addresses, Companies, Phones, etc). 
#### Create a new Faker
```go
    var f generator.Faker
    f = f.New(false) // False if we don't want the data to be unique
``` 
 
#### Misc Data (no language/locale)
Just basic data from native data types. 
```go
    var f generator.Faker
    f = f.New(false) 
    f.Misc.RandomIntBelow(400)   // Int below 400 
    f.Misc.Md5()                 // Random MD5 hashed string    
    f.Misc.UnixTime()            // Actual UNIX Timestamp
	// etc
``` 

### License
This code is free to use under the terms of the MIT license. See LICENSE.md for more information.