/// Represents a token in the alias definition syntax.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum Token {
    /// An unquoted identifier (e.g., `dk`, `git`)
    Identifier(String),
    /// A quoted string (e.g., `'git status'`)
    QuotedString(String),
    /// The equals sign separator
    Equals,
}

/// Lexer for tokenizing alias definition lines.
pub struct Lexer {
    input: Vec<char>,
    position: usize,
}

impl Lexer {
    pub fn new(input: &str) -> Self {
        Self {
            input: input.chars().collect(),
            position: 0,
        }
    }

    /// Tokenizes the entire input string.
    pub fn tokenize(&mut self) -> Vec<Token> {
        let mut tokens = Vec::new();

        while !self.is_at_end() {
            if let Some(token) = self.next_token() {
                tokens.push(token);
            }
        }

        tokens
    }

    fn next_token(&mut self) -> Option<Token> {
        self.skip_whitespace();

        if self.is_at_end() {
            return None;
        }

        let ch = self.current_char()?;

        match ch {
            '=' => {
                self.advance();
                Some(Token::Equals)
            }
            '\'' => Some(self.read_quoted_string()),
            _ => Some(self.read_identifier()),
        }
    }

    fn read_quoted_string(&mut self) -> Token {
        self.advance(); // Skip opening quote

        let mut content = String::new();
        let mut is_escaped = false;

        while !self.is_at_end() {
            let ch = self.current_char().unwrap();

            if is_escaped {
                content.push(ch);
                is_escaped = false;
            } else if ch == '\\' {
                is_escaped = true;
            } else if ch == '\'' {
                self.advance(); // Skip closing quote
                break;
            } else {
                content.push(ch);
            }

            self.advance();
        }

        Token::QuotedString(content)
    }

    fn read_identifier(&mut self) -> Token {
        let mut identifier = String::new();

        while !self.is_at_end() {
            let ch = self.current_char().unwrap();

            if ch == '=' || ch == '\'' || ch.is_whitespace() {
                break;
            }

            identifier.push(ch);
            self.advance();
        }

        Token::Identifier(identifier)
    }

    fn skip_whitespace(&mut self) {
        while !self.is_at_end() {
            if let Some(ch) = self.current_char() {
                if !ch.is_whitespace() {
                    break;
                }
                self.advance();
            }
        }
    }

    fn current_char(&self) -> Option<char> {
        self.input.get(self.position).copied()
    }

    fn advance(&mut self) {
        self.position += 1;
    }

    fn is_at_end(&self) -> bool {
        self.position >= self.input.len()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_tokenize_simple_alias() {
        let mut lexer = Lexer::new("dk=docker");
        let tokens = lexer.tokenize();

        assert_eq!(
            tokens,
            vec![
                Token::Identifier("dk".to_string()),
                Token::Equals,
                Token::Identifier("docker".to_string()),
            ]
        );
    }

    #[test]
    fn test_tokenize_quoted_expansion() {
        let mut lexer = Lexer::new("gb='git branch'");
        let tokens = lexer.tokenize();

        assert_eq!(
            tokens,
            vec![
                Token::Identifier("gb".to_string()),
                Token::Equals,
                Token::QuotedString("git branch".to_string()),
            ]
        );
    }

    #[test]
    fn test_tokenize_both_quoted() {
        let mut lexer = Lexer::new("'g cb'='git checkout -b'");
        let tokens = lexer.tokenize();

        assert_eq!(
            tokens,
            vec![
                Token::QuotedString("g cb".to_string()),
                Token::Equals,
                Token::QuotedString("git checkout -b".to_string()),
            ]
        );
    }

    #[test]
    fn test_tokenize_with_escape() {
        let mut lexer = Lexer::new("test='it\\'s working'");
        let tokens = lexer.tokenize();

        assert_eq!(
            tokens,
            vec![
                Token::Identifier("test".to_string()),
                Token::Equals,
                Token::QuotedString("it's working".to_string()),
            ]
        );
    }
}
