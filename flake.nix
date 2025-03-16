{
  description = "task-management-be";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-24.05-darwin";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in with pkgs; let
        nodeEnv = [
          nodejs_18
          corepack_18
        ];

        goEnv = [
          go
          gopls
          golangci-lint
        ];

        libraries = [
          docker-compose
          git
          jq
          pre-commit
          protobuf
          python3
          yq
        ];

        systemEnv = if (stdenv.isDarwin) then [darwin.apple_sdk_11_0.frameworks.Cocoa darwin.apple_sdk_11_0.frameworks.Security] else [systemd];

        defaultComposeFile = "docker-compose.yaml";
      in {
        devShell = pkgs.mkShell {
          buildInputs = nodeEnv ++ goEnv ++ libraries ++ systemEnv;

          COMPOSE_FILE = defaultComposeFile;
          ENV = "local";

          GORACE = "abort_on_error=1 halt_on_error=1";

          shellHook = ''
            pre-commit install --install-hooks -t pre-commit -t commit-msg
            export GOBIN=$(pwd)/dist/tools
            export CGO_ENABLED=1
            export PATH="$GOBIN:$PATH"
          '';
       };
      }
    );

}
