# concurrent-filesearch
simple project that demonstrate the mechanism of waitgroups in GoLang concorrent programs. the program is all about searching for specific file path in the whole driver. algorithm is quite simple, if the path contains the pattern, I check for if it is a file or a directory. if it was a file I simply add the string to a global slice as a matched result, if it was a directory I invoke an other thread to search all it's sub directories.
each search in each branch of directory will be handled by a thread in a concurrent way and recursive fashion.
to prevent the race condition, I used the mutex to syncronize the access to mathces slice.
