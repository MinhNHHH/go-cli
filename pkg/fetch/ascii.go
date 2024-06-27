package fetch

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type AsciiArt struct {
	Lines          []string
	MaxCleanLen    int
	MaxOriginalLen int
}

var Art = map[string]string{
	"linux": `            ${c1}.-/+oossssoo+/-.
        ${c1}":+ssssssssssssssssss+:"
      ${c1}-+ssssssssssssssssssyyssss+-
    ${c1}.ossssssssssssssssss${c7}dMMMNy${c1}sssso.
   ${c1}/sssssssssss${c7}hdmmNNmmyNMMMMh${c1}ssssss/
  ${c1}+sssssssss${c7}h${c1}myd${c7}MMMMMMMNdddd${c1}yssssssss+
 ${c1}/ssssssssh${c7}NMM${c1}Myh${c7}hyyyyhmNMMMNh${c1}ssssssss/
${c1}.ssssssss${c7}dMMMNh${c1}ssssssssss${c7}hNMMMd${c1}ssssssss.
${c1}+ssss${c7}hhhyNMMNy${c1}ssssssssssss${c7}yNMMM${c1}ysssssss+
${c1}ossy${c7}NMMMNyMMh${c1}sssssssssssssshmmmhssssssso
${c1}ossy${c7}NMMMNyMMh${c1}sssssssssssssshmmmhssssssso
${c1}+ssss${c7}hhhyNMMNy${c1}ssssssssssss${c7}yNMMM${c1}ysssssss+
${c1}.ssssssss${c7}dMMMNh${c1}ssssssssss${c7}hNMMM${c1}dssssssss.
 ${c1}/ssssssss${c7}hNMM${c1}Myh${c7}hyyyyhdNMMMNh${c1}ssssssss/
  ${c1}+sssssssss${c7}d${c1}myd${c7}MMMMMMMMddddy${c1}ssssssss+
   ${c1}/sssssssssssh${c7}dmNNNNmyNMMMMh${c1}ssssss/
    ${c1}.ossssssssssssssssss${c7}dMMMN$${c1}ysssso.
      ${c1}-+sssssssssssssssssyyyssss+-
        ${c1}":+ssssssssssssssssss+:"
            ${c1}.-/+oossssoo+/-.

`,
	"windows": `${c1}        ,.=:!!t3Z3z.,
       :tt:::tt333EE3
${c1}       Et:::ztt33EEEL${c2} @Ee.,      ..,
${c1}      ;tt:::tt333EE7${c2} ;EEEEEEttttt33#
${c1}     :Et:::zt333EEQ.${c2} $EEEEEttttt33QL
${c1}     it::::tt333EEF${c2} @EEEEEEttttt33F
${c1}    ;3=*^""""*4EEV${c2} :EEEEEEttttt33@.
${c3}    ,.=::::!t=., ${c1}"${c2} @EEEEEEtttz33QF
${c3}   ;::::::::zt33)${c2}   "4EEEtttji3P*
${c3}  :t::::::::tt33.${c4}:Z3z..${c2}  ""${c4} ,..g.    
${c3}  i::::::::zt33F${c4} AEEEtttt::::ztF
${c3} ;:::::::::t33V${c4} ;EEEttttt::::t3
${c3} E::::::::zt33L${c4} @EEEtttt::::z3F
${c3}{3=*^""""*4E3)${c4} ;EEEtttt:::::tZ"
${c3}             "${c4} :EEEEtttt::::z7
		${c4}"VEzjt:;;z>*"
		`,
	"darwin": `${c2}                    c.'
${c2}                 ,xNMM.
${c2}               .OMMMMo
${c2}               lMMM"
${c2}     .;loddo:. .olloddol;.
${c2}   cKMMMMMMMMMMNWMMMMMMMMMM0:
${c4} .KMMMMMMMMMMMMMMMMMMMMMMMWd.
${c4} XMMMMMMMMMMMMMMMMMMMMMMMX.
${c1};MMMMMMMMMMMMMMMMMMMMMMMM:
${c1}:MMMMMMMMMMMMMMMMMMMMMMMM:
${c1}.MMMMMMMMMMMMMMMMMMMMMMMMX.
${c1} kMMMMMMMMMMMMMMMMMMMMMMMMWd.
 ${c5}'XMMMMMMMMMMMMMMMMMMMMMMMMMMk
 ${c5}'XMMMMMMMMMMMMMMMMMMMMMMMMK.
  ${c3}kMMMMMMMMMMMMMMMMMMMMMMd
   ${c3};KMMMMMMMWXXWMMMMMMMk.
    ${c3}"cooc*"    "*coo'"`,
}

func NewAsciiArt(filePath string) (*AsciiArt, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	clean := math.MinInt
	original := math.MinInt
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		asc := scanner.Text()
		lines = append(lines, asc)

		original = Max(original, len(asc))
		for placeholder := range placeHolder {
			asc = strings.ReplaceAll(asc, placeholder, "")
		}
		clean = Max(len(asc), clean)
	}
	return &AsciiArt{
		Lines:          lines,
		MaxCleanLen:    clean,
		MaxOriginalLen: original,
	}, nil
}

func DefaultArt(sys string) *AsciiArt {
	scanner := bufio.NewScanner(strings.NewReader(Art[sys]))
	clean := math.MinInt
	original := math.MinInt
	lines := []string{}
	for scanner.Scan() {
		asc := scanner.Text()
		lines = append(lines, asc)

		original = Max(original, len(asc))
		for placeholder := range placeHolder {
			asc = strings.ReplaceAll(asc, placeholder, "")
		}
		clean = Max(len(asc), clean)
	}
	return &AsciiArt{
		Lines:          lines,
		MaxCleanLen:    clean,
		MaxOriginalLen: original,
	}
}
