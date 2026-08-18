package = "voxgig-sdk-art-institute-of-chicago"
version = "0.0.1-1"
source = {
  -- git+https (GitHub dropped git:// in 2022); pin the install to the release
  -- tag pushed by `make publish`, and point at the lua/ subdir of the monorepo.
  url = "git+https://github.com/voxgig-sdk/art-institute-of-chicago-sdk.git",
  tag = "lua/v0.0.1",
  dir = "art-institute-of-chicago-sdk/lua"
}
description = {
  summary = "Unofficial generated Lua SDK for the Art Institution of Chicago public API. Not affiliated with or endorsed by the upstream API provider.",
  homepage = "https://github.com/voxgig-sdk/art-institute-of-chicago-sdk",
  issues_url = "https://github.com/voxgig-sdk/art-institute-of-chicago-sdk/issues",
  license = "MIT",
  labels = { "voxgig", "sdk", "generated-sdk", "openapi", "api-client", "art-institute-of-chicago" }
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["art-institute-of-chicago_sdk"] = "art-institute-of-chicago_sdk.lua",
    ["config"] = "config.lua",
    ["config_shared"] = "config_shared.lua",
    ["features"] = "features.lua",
  }
}
