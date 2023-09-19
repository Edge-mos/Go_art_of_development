Построение бинарника происходит по схеме:
    go build -o {путь до папки, куда будет положен бинарник} {путь до файла-исходника}

Полные пути:
go build -o /Users/edge_mos/Desktop/goLang/go_lessons/art_of_development/build/primitive_types /Users/edge_mos/Desktop/goLang/go_lessons/art_of_development/1_primitive_types/primitive_types.go

Относительные пути из корневой папки(art_of_development):
go build -o build/primitive_types_four 1_primitive_types/primitive_types.go