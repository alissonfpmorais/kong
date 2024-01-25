{ pkgs ? import <nixpkgs> {
  config.allowUnfree = true;
} }:

with pkgs;

pkgs.mkShell {
  buildInputs = [
    delve
    go
    gopls
    go-tools
  ];
}
