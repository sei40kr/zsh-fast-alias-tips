/// Represents an alias definition with its name and expansion.
#[derive(Debug, Clone, PartialEq, Eq)]
pub struct AliasDefinition {
    /// The alias name (e.g., "gst")
    pub name: String,
    /// The expanded command (e.g., "git status")
    pub expansion: String,
}

impl AliasDefinition {
    pub fn new(name: String, expansion: String) -> Self {
        Self { name, expansion }
    }
}
