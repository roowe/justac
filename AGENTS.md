# Zig 本地资料查询

- 回答 Zig 语言、标准库、编译器实现或 API 问题时，优先查询本地 Zig 仓库：`/Users/luoliwei/github/zig`。
- 不要为 Zig 相关问题自动搜索互联网。只有用户明确要求联网查询时，才可以使用网络资料。
- 标准库以 `/Users/luoliwei/github/zig/lib/std` 中的源码和文档注释为准；语言参考资料优先查看 `/Users/luoliwei/github/zig/doc`。
- 优先使用 `rg` 搜索符号和调用位置，使用 `fd` 查找文件。例如：

  ```bash
  rg -n "pub fn print" /Users/luoliwei/github/zig/lib/std
  fd 'debug\.zig' /Users/luoliwei/github/zig/lib/std
  ```

- 回答时尽量给出本地源码的绝对路径和行号，方便直接定位。
- 如果本地仓库不存在、缺少目标内容，或版本与当前项目不匹配，应明确说明；不要自行改用网络搜索。
