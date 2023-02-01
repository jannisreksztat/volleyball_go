package latexDraw

const volley_template = `
\documentclass{article}
\usepackage{tikz}
\begin{document}

\begin{tikzpicture}

{{.TikzTemplate}}

\draw[fieldLine] (-{{.FeldLinie}}/1.9,0) -- ({{.FeldLinie}}/1.9,0);
{{range .VolleyField}}
\node[field={{.Size}}] ({{.Name}}) at ({{.XPos}},{{.YPos}}) { {{.Name}} };
{{end}}

{{range .Aktion}}
\draw[ass arrow] ({{.Startzone}}.center) -- ({{.Endzone}}.center) node [midway, right] { \{{.Rating}} };
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
        line width = 0.2mm,
    },
    ass arrow/.style={
        ->,
        red,
        line width = 0.4mm,
    },
    fieldLine/.style={
        black,
        line width = 0.5mm,
    },
}
`
