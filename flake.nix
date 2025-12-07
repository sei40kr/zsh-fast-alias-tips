{
  description = "Help remembering the aliases you defined once";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        zshrc = pkgs.writeText "zshrc" ''
          source ${pkgs.zinit}/share/zinit/zinit.zsh

          # Load the plugin from current directory
          zinit ice atclone'cargo build --release' atpull'%atclone'
          zinit load $PWD

          # Test aliases
          alias gst='git status'
          alias gco='git checkout'
          alias gcb='git checkout -b'

          echo "========================================"
          echo "zsh-fast-alias-tips test environment"
          echo "========================================"
          echo ""
          echo "Available test aliases:"
          echo "  gst -> git status"
          echo "  gco -> git checkout"
          echo "  gcb -> git checkout -b"
          echo ""
          echo "Try running:"
          echo "  git status"
          echo "  git checkout -b feature"
          echo ""
          echo "You should see alias tips appear!"
          echo "========================================"
        '';
      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            zsh
            zinit
            git

            cargo
            rustc
          ];

          shellHook = ''
            export ZDOTDIR=$(mktemp -d)
            ln -sf ${zshrc} $ZDOTDIR/.zshrc
            exec zsh
          '';
        };
      }
    );
}
