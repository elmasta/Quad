package piscine

import (
	"strconv"
)

func Quadchecker(str string) string {
	var stringTable []string
	temp := ""
	for _, v := range str {
		if v == '\n' {
			stringTable = append(stringTable, temp)
			temp = ""
		} else {
			temp += string(v)
		}
	}
	if len(stringTable) > 0 {
		compare := len(stringTable[0])
		isOk := true
		for _, v := range stringTable {
			if len(v) != compare {
				
				isOk = false
			}
		}
		if isOk {
			toReturn := ""
			largeur := len(stringTable[0])
			hauteur := len(stringTable)
			prob := false
			//potentiel quada
			if stringTable[0][0] == 'o' {
				for i, v := range stringTable {
					// première et dernière ligne
					if i == 0 || i == hauteur-1 {
						//scan de la chaine
						for subI, subV := range v {
							// coins du tableau
							if (subI == 0 || subI == largeur -1) && subV != 'o' {
								prob = true
							} else if (subI > 0 && subI < largeur -1) && subV != '-' {
								prob = true
							}
							if prob {
								break
							} else if i == hauteur-1 && subI == largeur -1 && !prob {
								toReturn = "[quadA] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
								return toReturn
							}
						}
					// milieux
					} else {
						//scan de la chaine
						for subI, subV := range v {
							if (subI == 0 || subI == largeur -1) && subV != '|' {
								prob = true
							} else if (subI > 0 && subI < largeur -1) && subV != ' ' {
								prob = true
							}
							if prob {
								break
							}
						}
					}
					if prob {
						break
					}
				}
			//potentiel quadb
			prob = false
			} else if stringTable[0][0] == '/' {
				for i, v := range stringTable {
					// première ligne
					if i == 0 {
						//scan de la première chaine
						for subI, subV := range v {
						// première ligne du tableau
							if (subI == 0 ) && subV != '/' {
								prob = true
							} else if subI != 0 && subI == largeur -1 && subV != '\\' {
								prob = true
							} else if (subI > 0 && subI < largeur -1) && subV != '*' {
								prob = true
							}
							if prob {
								break
							} else if i == hauteur-1 && subI == largeur -1 && !prob {
								toReturn = "[quadB] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
								return toReturn
							}
						}
					// dernière ligne
					} else if i == hauteur-1 {
						for subI, subV := range v {
							if subI == 0 && subV != '\\' {
								prob = true
							} else if subI != 0 && subI == largeur -1 && subV != '/' {
								prob = true
							} else if subI > 0 && subI < largeur -1 && subV != '*' {
								prob = true
							}
							if prob {	
								break
							} else if i == hauteur-1 && subI == largeur -1 && !prob {
								toReturn = "[quadB] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
								return toReturn
							}
						}
					} else {
						//scan de la chaine milieu
						for subI, subV := range v {
							if (subI == 0 || subI == largeur -1) && subV != '*' {
								prob = true
							} else if (subI > 0 && subI < largeur -1) && subV != ' ' {
								prob = true
							}
							if prob {
								break
							}
						}
					}
					if prob{
						break
					}
				}
			} else {
				pQuadC := false
				pQuadD := false
				pQuadE := false
				if len(stringTable[0]) == 1 && len(stringTable) == 1 {
					return "[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]"
				} else {
					//check quad c
					for i, v := range stringTable {
						// première ligne
						if i == 0 {
							//scan de la chaine
							for subI, subV := range v {
								// coins du tableau
								if (subI == 0 || subI == largeur -1) && subV != 'A' {
									
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != 'B' {
									prob = true
								}
								if prob {
									break
								} else if i == hauteur-1 && subI == largeur -1 && !prob {
									pQuadC = true
								}
							}
						// fin
						} else if i == hauteur-1 {
							// coins du tableau
							for subI, subV := range v {
								if (subI == 0 || subI == largeur -1) && subV != 'C' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != 'B' {
									prob = true
								}
								if prob {
									break
								} else if i == hauteur-1 && subI == largeur -1 && !prob {
									pQuadC = true
								}
							}
						// milieux
						} else {
							//scan de la chaine
							for subI, subV := range v {
								if (subI == 0 || subI == largeur -1) && subV != 'B' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != ' ' {
									prob = true
								}
								if prob {
									break
								}
							}
						}
						if prob {
							break
						}
					}
					//check quad d
					prob = false
					for i, v := range stringTable {
						// première ligne
						if i == 0 || i == hauteur-1 {
							//scan de la chaine
							for subI, subV := range v {
								// coins du tableau
								if subI == 0 && subV != 'A' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != 'B' {
									prob = true
								} else if largeur != 1 && subI == largeur -1 && subV != 'C' {
									prob = true
								}
								if prob {
									break
								} else if i == hauteur-1 && subI == largeur -1 && !prob {
									pQuadD = true
								}
							}
						// milieux
						} else {
							//scan de la chaine
							for subI, subV := range v {
								if (subI == 0 || subI == largeur -1) && subV != 'B' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != ' ' {
									prob = true
								}
								if prob {
									break
								}
							}
						}
						if prob {
							break
						}
					}
					//check quad e
					prob = false
					for i, v := range stringTable {
						// première ligne
						if i == 0 {
							//scan de la chaine
							for subI, subV := range v {
								// coins du tableau
								if subI == 0 && subV != 'A' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != 'B' {
									prob = true
								} else if largeur != 1 && subI == largeur -1 && subV != 'C' {
									prob = true
								}
								if prob {
									break
								} else if i == hauteur-1 && subI == largeur -1 && !prob {
									pQuadE = true
								}
							}
						// fin
						} else if i == hauteur-1 {
							// coins du tableau
							for subI, subV := range v {
								if subI == 0 && subV != 'C' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != 'B' {
									prob = true
								} else if subI != 0 && subI == largeur -1 && subV != 'A' {
									prob = true
								}
								if prob {
									break
								} else if i == hauteur-1 && subI == largeur -1 && !prob {
									pQuadE = true
								}
							}
						// milieux
						} else {
							//scan de la chaine
							for subI, subV := range v {
								if (subI == 0 || subI == largeur -1) && subV != 'B' {
									prob = true
								} else if (subI > 0 && subI < largeur -1) && subV != ' ' {
									prob = true
								}
								if prob {
									break
								}
							}
						}
						if prob {
							break
						}
					}
				}
				if !pQuadC && !pQuadD && !pQuadE {
					return "Not a quad function"
				} else {
					if pQuadC {
						toReturn = "[quadC] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
					}
					if pQuadD {
						if len(toReturn) > 0 {
							toReturn += " || "
						}
						toReturn += "[quadD] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
					}
					if pQuadE {
						if len(toReturn) > 0 {
							toReturn += " || "
						}
						toReturn += "[quadE] [" + strconv.Itoa(largeur) + "] [" + strconv.Itoa(hauteur) + "]"
					}
					return toReturn
				}
			}
			if prob {
				return "Not a quad function"
			} else {
				return toReturn
			}
		}
	}
	return "Not a quad function"
}
