// mayhem/lsan_off.c — turn LeakSanitizer off at build time (SPEC §6.2 item 15). ASan stays on;
// only leak detection, which this fleet does not fuzz for, is disabled.
int __lsan_is_turned_off(void) { return 1; }
