use crate::model::AliasDefinition;

/// The result of matching a command against alias definitions.
#[derive(Debug)]
pub struct MatchResult<'a> {
    /// The matched alias definition
    pub definition: &'a AliasDefinition,
    /// Whether the match is a full match (true) or partial match (false)
    pub is_full_match: bool,
}

/// Finds the best matching alias for a given command.
///
/// This function implements a recursive matching algorithm:
/// 1. Sorts aliases by expansion length (longest first)
/// 2. Finds the longest matching alias
/// 3. Substitutes the matched alias and repeats
/// 4. Returns the final matched alias
///
/// # Examples
///
/// ```
/// # use alias_matcher::model::AliasDefinition;
/// # use alias_matcher::matcher::find_best_match;
/// let defs = vec![
///     AliasDefinition::new("dk".to_string(), "docker".to_string()),
///     AliasDefinition::new("gst".to_string(), "git status".to_string()),
/// ];
///
/// let result = find_best_match(&defs, "docker");
/// assert!(result.is_some());
/// assert_eq!(result.unwrap().definition.name, "dk");
/// ```
pub fn find_best_match<'a>(
    definitions: &'a [AliasDefinition],
    command: &str,
) -> Option<MatchResult<'a>> {
    let mut current_command = command.to_string();
    let mut candidate: Option<&'a AliasDefinition> = None;
    let mut is_full_match = false;

    loop {
        let matched = find_longest_match(definitions, &current_command);

        match matched {
            Some((def, is_full)) => {
                current_command =
                    format!("{}{}", def.name, &current_command[def.expansion.len()..]);
                candidate = Some(def);
                is_full_match = is_full;
            }
            None => break,
        }
    }

    candidate.map(|def| MatchResult {
        definition: def,
        is_full_match,
    })
}

/// Finds the longest matching alias for the given command.
fn find_longest_match<'a>(
    definitions: &'a [AliasDefinition],
    command: &str,
) -> Option<(&'a AliasDefinition, bool)> {
    let mut best_match: Option<(&'a AliasDefinition, bool)> = None;
    let mut best_length = 0;

    for def in definitions {
        if command == def.expansion && def.expansion.len() > best_length {
            best_match = Some((def, true));
            best_length = def.expansion.len();
        } else if command.starts_with(&def.expansion) && def.expansion.len() > best_length {
            best_match = Some((def, false));
            best_length = def.expansion.len();
        }
    }

    best_match
}

#[cfg(test)]
mod tests {
    use super::*;

    fn create_test_definitions() -> Vec<AliasDefinition> {
        vec![
            AliasDefinition::new("dk".to_string(), "docker".to_string()),
            AliasDefinition::new("gb".to_string(), "git branch".to_string()),
            AliasDefinition::new("gco".to_string(), "git checkout".to_string()),
            AliasDefinition::new("gcb".to_string(), "git checkout -b".to_string()),
            AliasDefinition::new("ls".to_string(), "ls -G".to_string()),
            AliasDefinition::new("ll".to_string(), "ls -lh".to_string()),
        ]
    }

    // Normal cases
    #[test]
    fn test_match_single_token() {
        let defs = create_test_definitions();
        let result = find_best_match(&defs, "docker");
        assert!(result.is_some());
        assert_eq!(result.unwrap().definition.name, "dk");
    }

    #[test]
    fn test_match_multiple_tokens() {
        let defs = create_test_definitions();
        let result = find_best_match(&defs, "git branch");
        assert!(result.is_some());
        assert_eq!(result.unwrap().definition.name, "gb");
    }

    #[test]
    fn test_match_prefers_longest() {
        let defs = create_test_definitions();
        let result = find_best_match(&defs, "git checkout -b");
        assert!(result.is_some());
        assert_eq!(result.unwrap().definition.name, "gcb");
    }

    #[test]
    fn test_match_recursive() {
        let defs = create_test_definitions();
        let result = find_best_match(&defs, "ls -G -lh");
        assert!(result.is_some());
        assert_eq!(result.unwrap().definition.name, "ll");
    }

    #[test]
    fn test_match_no_matches() {
        let defs = create_test_definitions();
        let result = find_best_match(&defs, "cd ..");
        assert!(result.is_none());
    }
}
