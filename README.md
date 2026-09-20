# Evaluation tableaux en Go

Ce projet est un exercice d'introduction au langage Go. Il permet de manipuler
un tableau de soldats, d'effectuer quelques calculs et de simuler plusieurs
attaques ennemies.

## Lancer le programme

Il faut avoir Go installé sur l'ordinateur. Depuis le dossier du projet, lancer
la commande suivante :

```bash
go run .
```

Pour vérifier que le code compile :

```bash
go test ./...
```

Le fichier `go.mod` indique que le dossier est un module Go.

## Organisation des données

Chaque soldat est représenté par la structure `Soldat` :

```go
type Soldat struct {
	nom     string
	vie     int
	attaque int
}
```

La structure contient trois informations :

- `nom` : le nom du soldat ;
- `vie` : ses points de vie ;
- `attaque` : sa puissance d'attaque.

L'équipe est un tableau de six soldats :

```go
var equipe = [6]Soldat{...}
```

Le nombre `6` signifie que le tableau contient exactement six éléments.

## Fonctions du programme

### `afficherEquipe`

Cette fonction parcourt le tableau avec une boucle `for` et affiche le nom,
l'attaque et les points de vie de chaque soldat.

### `trouverPlusDeVie`

Cette fonction recherche le soldat qui possède le plus de points de vie.
Elle commence par considérer le premier soldat comme le meilleur, puis compare
les autres soldats un par un.

### `trouverPlusDAttaque`

Cette fonction utilise le même principe que `trouverPlusDeVie`, mais elle
compare la valeur `attaque`.

### `calculerVieMoyenne`

La fonction additionne les points de vie de tous les soldats, puis divise le
total par le nombre de soldats. Elle retourne un `float64` pour pouvoir
afficher une moyenne avec des décimales.

### `compterFaibles`

Cette fonction compte les soldats dont les points de vie sont inférieurs à
800. Un compteur est augmenté chaque fois que la condition est vraie.

### `attaquerEquipe`

Cette fonction applique les mêmes dégâts à tous les soldats :

```go
func attaquerEquipe(equipe *[6]Soldat, degats int)
```

Le symbole `*` signifie que la fonction reçoit l'adresse du tableau. Elle peut
donc modifier directement les points de vie de l'équipe.

Les points de vie ne peuvent pas devenir négatifs. Si une attaque retire plus
de points de vie qu'un soldat n'en possède, ses points de vie sont ramenés à
`0`.

### `afficherEtat`

Cette fonction affiche l'état de chaque soldat après une attaque. Un soldat
avec `0` point de vie est affiché comme `KO`.

### `compterVivants`

Cette fonction compte les soldats encore vivants en utilisant la récursivité :

```go
func compterVivants(equipe [6]Soldat, index int) int
```

Un soldat est vivant lorsque `vie > 0`.

La fonction ne contient pas de boucle. Elle fonctionne en deux étapes :

1. Elle regarde le soldat situé à la position `index`.
2. Elle s'appelle elle-même avec `index + 1` pour regarder le soldat suivant.

Le cas d'arrêt est atteint lorsque `index` arrive à la fin du tableau :

```go
if index == len(equipe) {
	return 0
}
```

À ce moment-là, il n'y a plus de soldat à vérifier. La fonction retourne `0`,
puis les appels précédents additionnent les soldats vivants.

Par exemple, pour une équipe contenant trois soldats vivants, le résultat est
`3`.

## Déroulement du programme

Le programme suit les étapes suivantes :

1. Afficher l'équipe de départ.
2. Afficher le soldat avec le plus de vie.
3. Afficher le soldat avec la plus grande attaque.
4. Calculer et afficher la vie moyenne.
5. Compter les soldats ayant moins de 800 points de vie.
6. Demander le nombre d'attaques ennemies.
7. Pour chaque attaque :
   - demander le nombre de dégâts ;
   - appeler `attaquerEquipe` ;
   - afficher l'état de l'équipe ;
   - compter les soldats vivants avec `compterVivants`.

## Exemple

Une exécution peut ressembler à ceci :

```text
=== BATAILLE ===

Nombre d'attaques ennemies : 2

Attaque 1 : 200

=== APRÈS L'ATTAQUE 1 ===

Arthas : 1000 PV
Kael : 650 PV
Thrall : 1300 PV
Sylvanas : 500 PV
Garrosh : 800 PV
Jaina : 300 PV
Soldats vivants : 6

Attaque 2 : 350

=== APRÈS L'ATTAQUE 2 ===

Arthas : 650 PV
Kael : 300 PV
Thrall : 950 PV
Sylvanas : 150 PV
Garrosh : 450 PV
Jaina : KO
Soldats vivants : 5
```