![entête](https://play.mcdueltown.fr/CSS/img/gt/entete.png)

# Groupie Tracker

Ce site vous affiche une liste d'artiste.

Redif des lives :

[🔴 Groupie Tracker, projet Golang/Html/CSS/JS (replay live partie 1) - YouTube](https://youtu.be/9474gO-5lRo)

[🔴 Groupie Tracker, projet Golang/Html/CSS/JS (replay live partie 2) - YouTube](https://youtu.be/x_gUDE3kgdg)

## 1. Installation

Le projet est hébergé sur le github de [Robin Guerard]([RobinGuerard21/GroupieTracker · GitHub](https://github.com/RobinGuerard21/GroupieTracker.git)).

Il vous suffit d'avoir les accès à ce répertoire et de le cloner grâce à git.

```bash
git https://github.com/RobinGuerard21/GroupieTracker.git
```

## 2. Le Lancement

Pour lancer le server il vous suffit d'exécuter le dossier dans lequel ce trouve votre main.go ou d'executer le package `main`.

Par défaut le projet sera hébergé sur l'ip 127.0.0.1 soit le localhost et le port 800.

Pour changer l'ip est le port rendez-vous dans le fichier main.go à la ligne 28.

```
golang
http.ListenAndServe("127.0.0.1:800", nil)
```

## 3. Se Rendre sur le Site

Il vous suffit d'aller à l'addresse : [localhost:800/index](http://localhost:800/index)

### 3.1 Page d'Accueil

![index](https://play.mcdueltown.fr/CSS/img/gt/gtindex.PNG)

Cette page vous présentera le site et son but. Vous pourrez aussi vous servir pour visiter notre magnifique site.

### 3.2 Page des Artistes

![header](https://play.mcdueltown.fr/CSS/img/gt/artisteheader.PNG)

Les bouton à droite vous permettrons de changer la disposition des artiste entre le mode bulle (par défaut), le mode ligne et le mode carte.

A gauche vous touverez une barre de recherche faite par nos soins qui vous permettra de filtrer les groupes/artistes par: leur Nom, leur Date de Création, leurs Membres, la date de sortie de leur premier album ou encore leur nombre de membres.

Lorsque vous clickez sur un artiste ou un groupe une popup s'ouvrira pour vous exposer plus de détails sur ce dernier. De plus en cliquant sur `See on map` vous serez redirigé sur la page [Map](https://git.ytrack.learn.ynov.com/RGUERARD/groupietracker#4-page-map) avec le filtre sur les lieux où se trouve leur futur concerts.

![popup](https://play.mcdueltown.fr/CSS/img/gt/artistedesc.PNG)

#### 3.2.a Mode Bulle

![Bulle](https://play.mcdueltown.fr/CSS/img/gt/artistebubble.PNG)

Est un mode d'affichage amusant que nous avons de choisit d'afficher par défaut passez votre souris au dessus d'une bulle et vous découvrirez un univers.

#### 3.2.b Mode Ligne

![Ligne](https://play.mcdueltown.fr/CSS/img/gt/artisterow.PNG)

Est un mode d'affichage classique mais tout de même revisité à notre goût.

#### 3.2.c Mode Carte

![Carte](https://play.mcdueltown.fr/CSS/img/gt/artistecard.PNG)

Comme ce sont des groupes/chanteur nous avons imaginé faire une disposition comme si c'était un lecteur mp4 que vous teniez dans la main.

## 4. Page Map

![Map](https://play.mcdueltown.fr/CSS/img/gt/gtmap.PNG)

Ici vous pourrez trouvez le lieu où se tiendra les futurs concerts de vos artiste préféré.

Lorsque vous clickez sur un point vous obtiendrez le nom de l'artiste ou du groupe ainsi que la date de sa représentation comme montré ci-dessous:

![MapDesc](https://play.mcdueltown.fr/CSS/img/gt/gtmapdesc.PNG)

Et si vous avez le paramettre d'url `artist` vous ne verrez que les emplacements pour l'artiste selectionné. Lorsque vous êtes redirigé sur la page grâce a la redirection de la popup de la page artiste ce paramètre sera automatiquement défini.

Exmeple de la redirection depuis la popup de Queen:

![MapQueen](https://play.mcdueltown.fr/CSS/img/gt/gtmapqueen.PNG)