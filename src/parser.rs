use crate::lexer::Token;
use crate::model::AliasDefinition;

/// Error type for parsing failures.
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum ParseError {
    /// Expected an equals sign but found something else
    ExpectedEquals,
    /// Expected a name (identifier or quoted string) but found something else
    ExpectedName,
    /// Expected an expansion (identifier or quoted string) but found something else
    ExpectedExpansion,
    /// Unexpected end of input
    UnexpectedEndOfInput,
}

/// Parser for alias definitions.
///
/// Grammar:
/// ```text
/// alias_def := name '=' expansion
/// name      := Identifier | QuotedString
/// expansion := Identifier | QuotedString
/// ```
pub struct Parser {
    tokens: Vec<Token>,
    position: usize,
}

impl Parser {
    pub fn new(tokens: Vec<Token>) -> Self {
        Self {
            tokens,
            position: 0,
        }
    }

    /// Parses an alias definition from the token stream.
    pub fn parse(&mut self) -> Result<AliasDefinition, ParseError> {
        let name = self.parse_name()?;
        self.expect_equals()?;
        let expansion = self.parse_expansion()?;

        Ok(AliasDefinition::new(name, expansion))
    }

    fn parse_name(&mut self) -> Result<String, ParseError> {
        match self.current_token() {
            Some(Token::Identifier(s)) | Some(Token::QuotedString(s)) => {
                let name = s.clone();
                self.advance();
                Ok(name)
            }
            Some(_) => Err(ParseError::ExpectedName),
            None => Err(ParseError::UnexpectedEndOfInput),
        }
    }

    fn expect_equals(&mut self) -> Result<(), ParseError> {
        match self.current_token() {
            Some(Token::Equals) => {
                self.advance();
                Ok(())
            }
            Some(_) => Err(ParseError::ExpectedEquals),
            None => Err(ParseError::UnexpectedEndOfInput),
        }
    }

    fn parse_expansion(&mut self) -> Result<String, ParseError> {
        match self.current_token() {
            Some(Token::Identifier(s)) | Some(Token::QuotedString(s)) => {
                let expansion = s.clone();
                self.advance();
                Ok(expansion)
            }
            Some(_) => Err(ParseError::ExpectedExpansion),
            None => Err(ParseError::UnexpectedEndOfInput),
        }
    }

    fn current_token(&self) -> Option<&Token> {
        self.tokens.get(self.position)
    }

    fn advance(&mut self) {
        self.position += 1;
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::lexer::Token;

    // Normal cases
    #[test]
    fn test_parser_success() {
        let tokens = vec![
            Token::Identifier("dk".to_string()),
            Token::Equals,
            Token::Identifier("docker".to_string()),
        ];
        let mut parser = Parser::new(tokens);
        let def = parser.parse().unwrap();
        assert_eq!(def.name, "dk");
        assert_eq!(def.expansion, "docker");
    }

    #[test]
    fn test_parser_with_quoted_strings() {
        let tokens = vec![
            Token::QuotedString("g cb".to_string()),
            Token::Equals,
            Token::QuotedString("git checkout -b".to_string()),
        ];
        let mut parser = Parser::new(tokens);
        let def = parser.parse().unwrap();
        assert_eq!(def.name, "g cb");
        assert_eq!(def.expansion, "git checkout -b");
    }

    // Semi-normal cases
    #[test]
    fn test_parser_expected_equals() {
        let tokens = vec![
            Token::Identifier("dk".to_string()),
            Token::Identifier("docker".to_string()),
        ];
        let mut parser = Parser::new(tokens);
        let result = parser.parse();
        assert_eq!(result, Err(ParseError::ExpectedEquals));
    }

    #[test]
    fn test_parser_expected_name() {
        let tokens = vec![Token::Equals, Token::Identifier("docker".to_string())];
        let mut parser = Parser::new(tokens);
        let result = parser.parse();
        assert_eq!(result, Err(ParseError::ExpectedName));
    }

    #[test]
    fn test_parser_expected_expansion() {
        let tokens = vec![Token::Identifier("dk".to_string()), Token::Equals];
        let mut parser = Parser::new(tokens);
        let result = parser.parse();
        assert_eq!(result, Err(ParseError::UnexpectedEndOfInput));
    }

    #[test]
    fn test_parser_unexpected_end_of_input_at_name() {
        let tokens = vec![];
        let mut parser = Parser::new(tokens);
        let result = parser.parse();
        assert_eq!(result, Err(ParseError::UnexpectedEndOfInput));
    }

    #[test]
    fn test_parser_unexpected_end_of_input_at_equals() {
        let tokens = vec![Token::Identifier("dk".to_string())];
        let mut parser = Parser::new(tokens);
        let result = parser.parse();
        assert_eq!(result, Err(ParseError::UnexpectedEndOfInput));
    }
}
