package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Structure de données pour le jeu
type JeuPendu struct {
	Mot            string
	MotADeviner    string
	TentativesRest int
	Etapes         []string
	LettresDejaDev []rune
}

func main() {
	jeu := JeuPendu{
		TentativesRest: 10,
	}
	jeu.Etapes = chargerEtapes("hangman.txt")
	jeu.MotADeviner = choisirMot("words.txt")
	jeu.Mot = revelerLettresAleatoires(jeu.MotADeviner)

	lecteur := bufio.NewReader(os.Stdin)

	for jeu.TentativesRest > 0 {
		fmt.Printf("Mot actuel: %s\n", jeu.Mot)
		fmt.Print("Entrez une lettre ou un mot (ou 'quit' pour sortir): ")
		entree, _ := lecteur.ReadString('\n')
		entree = strings.TrimSpace(entree)

		if strings.ToLower(entree) == "quit" {
			fmt.Println("Vous avez quitté le jeu.")
			return
		}

		if len(entree) == 1 {
			devinette := rune(entree[0])

			if contient(jeu.LettresDejaDev, devinette) {
				fmt.Println("Cette lettre a déjà été devinée.")
				continue
			}

			jeu.LettresDejaDev = append(jeu.LettresDejaDev, devinette)

			if strings.ContainsRune(jeu.MotADeviner, devinette) {
				jeu.Mot = revelerLettres(jeu.MotADeviner, jeu.Mot, devinette)
			} else {
				jeu.TentativesRest--
				fmt.Printf("Incorrect ! Tentatives restantes: %d\n", jeu.TentativesRest)
				fmt.Println(jeu.Etapes[10-jeu.TentativesRest])
			}
		} else {
			if entree == jeu.MotADeviner {
				fmt.Println("Bravo ! Vous avez trouvé le mot:", jeu.MotADeviner)
				return
			} else {
				jeu.TentativesRest -= 2
				fmt.Printf("Mauvais mot ! Tentatives restantes: %d\n", jeu.TentativesRest)
				fmt.Println(jeu.Etapes[10-jeu.TentativesRest])
			}
		}

		if jeu.Mot == jeu.MotADeviner {
			fmt.Println("Bien joué ! Vous avez deviné le mot:", jeu.MotADeviner)
			return
		}
	}

	fmt.Println("Jeu terminé ! Le mot était:", jeu.MotADeviner)
	fmt.Println(jeu.Etapes[len(jeu.Etapes)-1])
}

func chargerEtapes(nomFichier string) []string {
	contenu, err := ioutil.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier:", err)
		os.Exit(1)
	}
	etapes := strings.Split(string(contenu), "=========\n")
	for i := range etapes {
		etapes[i] = strings.TrimSpace(etapes[i]) + "\n========="
	}
	return etapes
}

func choisirMot(nomFichier string) string {
	contenu, err := ioutil.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier:", err)
		os.Exit(1)
	}
	mots := strings.Split(string(contenu), "\n")
	rand.Seed(time.Now().UnixNano())
	return mots[rand.Intn(len(mots))]
}

func revelerLettresAleatoires(mot string) string {
	revele := make([]rune, len(mot))
	for i := range revele {
		revele[i] = '_'
	}
	n := len(mot) / 2
	for i := 0; i < n; i++ {
		index := rand.Intn(len(mot))
		revele[index] = rune(mot[index])
	}
	return string(revele)
}

func revelerLettres(aDeviner, courant string, devinette rune) string {
	resultat := []rune(courant)
	for i, c := range aDeviner {
		if c == devinette {
			resultat[i] = c
		}
	}
	return string(resultat)
}

func contient(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
