# Misc
Since version 1.0.0

This Generator will provide a full set of miscellaneous data. From basic data types (string, int)
to more complex data like hashes, lists, etc.

This is the only generator that does not require a locale configuration.

```go
    MiscGenerator.RandomInt()           // 65
    MiscGenerator.RandomFloat32()       // 0.2051103
    MiscGenerator.RandomFloat64()       // 0.6046602879796196
    MiscGenerator.RandomIntBelow(10)    // 4
    MiscGenerator.UnixTime()            // 1257894000
    MiscGenerator.Md5()                 // e5ac1579769873fe68790e5149a05779
    MiscGenerator.Sha1()                // 4e227d2f3abeaadd6e80f9d14db231d1bd4c329b
    MiscGenerator.Sha256()              // 2a92db8d05aaaa0c9aac7e68cd7300dc8c6fb6abdd50595ebd16f521148a120f
```

