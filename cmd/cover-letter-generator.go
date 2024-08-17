package cmd

import (
	"fmt"
	"strings"
)

var DEFAULT_COVER_LETTER_TEMPLATE = `
Pouya Tavakoli
405-887-7447 | 1pouya.tavakoli@gmail.com
To {company}
I am interested in the {position} position at {company}. I found this position on
your career website and after reading the job description I am confident that I will be a good fit
for the role.
Through my professional career I’ve learned a lot on managing infrastructure at scale. I
currently work with 50+ clusters per environment with each one having around 100-200 pods.
With the largest cluster having around ~5k pods. I’m also part of an on call rotation where I fix
any issues that may arise in production.
I also have experience with handling and getting up to speed with legacy platforms and
spearheading initiatives to migrate customers off of them.
Pouya Tavakoli
`

// generatePDF(content string, filename string) error {

// }

func GenerateCoverLetter(company string, position string) {
	replacer := strings.NewReplacer(
		"{company}", company,
		"{position}", position,
	)
	final_cover_letter := replacer.Replace(DEFAULT_COVER_LETTER_TEMPLATE)
	// filename := company + ".pdf"
	fmt.Println(final_cover_letter)

}
