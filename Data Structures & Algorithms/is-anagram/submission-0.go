func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var counts [26]int

    for _, sym_t := range t {
        counts[sym_t-'a']++
    }
    for _, sym_s := range s {
        counts[sym_s-'a']--
        if counts[sym_s-'a'] < 0 {
            return false
        }
    }
    return true
}