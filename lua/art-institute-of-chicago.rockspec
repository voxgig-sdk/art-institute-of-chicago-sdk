package = "voxgig-sdk-art-institute-of-chicago"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/art-institute-of-chicago-sdk.git"
}
description = {
  summary = "ArtInstituteOfChicago SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["art-institute-of-chicago_sdk"] = "art-institute-of-chicago_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
