package latexDraw

const volley_template = `
\documentclass{article}
\usepackage{tikz}
\begin{document}

\begin{tikzpicture}

{{.TikzTemplate}}

{{range .VolleyField}}
\node[field={{.Size}}] ({{.Name}}) at ({{.XPos}},{{.YPos}}) { {{.Name}} };
{{end}}

\end{tikzpicture}
\end{document}
`
const tikz_template = `
\tikzset{
    field/.style ={
        rectangle,
        draw,
        minimum width = #1*1cm, 
        minimum height = #1*1cm, 
        line width = 0.5mm
    },
    ass arrow/.style={
        ->,
        red,
        line width = 0.5mm
    },
}
`
