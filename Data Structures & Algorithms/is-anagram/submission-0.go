func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    s_map := make(map[rune]int)

    for _, s_char := range s {
        _, exists := s_map[s_char]
        if exists {
            s_map[s_char] = s_map[s_char] + 1
        } else {
            s_map[s_char] = 1
        }
    }

    for _, t_char := range t {
        _, exists := s_map[t_char]
        if exists {
            s_map[t_char] = s_map[t_char] - 1
        } else {
            return false
        }
    }

    for _, value := range s_map {
        if value != 0 {
            return false
        }
    }
    return true
}
