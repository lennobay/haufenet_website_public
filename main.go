package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

var data_projects_index = ``
var main_page_top = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dauguls | Start</title>
	<link rel="stylesheet" href="/style.css">
	<link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
<link rel="manifest" href="/site.webmanifest">
</head>
<body>
    <content>
        <header>
            <p style="font-weight: bolder;" onclick="window.location.href='/'">Dauguls</p>
            <p style="font-weight: 500;" onclick="window.location.href='/projects'">Projects</p>
            <p style="font-weight: 500;" onclick="window.location.href='/articles'">Articles</p>
            <p style="font-weight: 500;" onclick="window.location.href='/about'">About</p>
        </header>
        

        <hero>
            <h1>Hi there!</h1>
            <p>I am an amature programmer and student in high school.</p>
            <button onclick="window.location.href='/#projects'">Look at my Projects</button>
        </hero>

        <projects id="projects">
            <projects-left>
                <h1>My Projects</h1>
                <p>Creating projects just for fun when ever I need or want something</p>
                <button onclick="window.location.href='/projects'">More Projects</button>
            </projects-left>
			<projects-right>`
var main_page_middle = `<button onclick="window.location.href='/projects'">More Projects</button>
            </projects-right>
        </projects>
        <articles>
            <articles-left>
                <h1>My Articles</h1>
                <p>Writing about expieriences and missing guides on the Internet</p>
            </articles-left>
            <articles-right>`
var main_page_bottom = `
			<button onclick="window.location.href='/articles'">Other Articles</button>
            </articles-right>
            </articles>


        <footer>
            <a href='/imprint_privacy'>Imprint and Privacy</a><p>&copy; dauguls.de</p>
        </footer>

    </content>
</body>
</html>`
var article_project_page_top = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dauguls</title>
	<link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
<link rel="manifest" href="/site.webmanifest">
    <link rel="stylesheet" href="/style.css">
</head>

<body>
    <content>
        <header>
            <p style="font-weight: bolder;" onclick="window.location.href='/'">Dauguls</p>
            <p style="font-weight: 500;" onclick="window.location.href='/projects'">Projects</p>
            <p style="font-weight: 500;" onclick="window.location.href='/articles'">Articles</p>
            <p style="font-weight: 500;" onclick="window.location.href='/about'">About</p>
        </header>
		<writing>
		<writing-content>
`
var article_project_page_bottom = ` </writing></writing-content>  <footer>
            <a href='/imprint_privacy'>Imprint and Privacy</a>
            <p>&copy; dauguls.de</p>
        </footer>

    </content>
</body>

</html>`
var main_page_articles_element = ""
var main_page_projects_element = ""
var data_articles_index = ""

func main() {
	CreateMainPage()
	CreateArticleMainPage()
	CreateArticlePages()
	CreateProjectPages()
	CreateProjectMainPage()
	readFileOthers("imprint_privacy.md")
	readFileOthers("about.md")
	fileserver := http.FileServer(http.Dir("./httpdocs"))

	http.Handle("/", fileserver)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
func CreateProjectPages() {
	err := os.Mkdir("./httpdocs/", 0755)
	err = os.Mkdir("./httpdocs/projects", 0755)
	entries, err := os.ReadDir("./projects")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles("./articles/", e.Name(), "ProjectsSinglePages")
	}
}
func CreateArticlePages() {
	err := os.Mkdir("./httpdocs/", 0755)
	err = os.Mkdir("./httpdocs/articles", 0755)
	entries, err := os.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles("./articles/", e.Name(), "ArticleSinglePages")
	}
}
func CreateArticleMainPage() {
	err := os.Mkdir("./httpdocs/articles", 0755)
	entries, err := os.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles("./articles/", e.Name(), "ArticleMainPage")
	}
	err = os.WriteFile("./httpdocs/articles/index.html", []byte(article_project_page_top+"<h1>Articles</h1>"+data_articles_index+article_project_page_bottom), 0755)
	if err != nil {
		panic(err)
	}

}
func CreateProjectMainPage() {
	err := os.Mkdir("./httpdocs/projects", 0755)
	entries, err := os.ReadDir("./projects")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles("./projects/", e.Name(), "ProjectMainPage")
	}
	err = os.WriteFile("./httpdocs/projects/index.html", []byte(article_project_page_top+"<h1>My Projects</h1>"+data_projects_index+article_project_page_bottom), 0755)
	if err != nil {
		panic(err)
	}

}
func CreateMainPage() {
	err := os.Mkdir("./httpdocs/", 0755)
	entries, err := os.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		analayzeFiles("./articles/", e.Name(), "MainPageArticles")
	}
	entries2, err := os.ReadDir("./projects")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries2 {
		analayzeFiles("./projects/", e.Name(), "MainPageProjects")
	}
	err = os.WriteFile("./httpdocs/index.html", []byte(main_page_top+main_page_projects_element+main_page_middle+main_page_articles_element+main_page_bottom), 0755)
	if err != nil {
		panic(err)
	}

}

func analayzeFiles(direc string, filename string, action string) {
	file, err := os.Open(direc + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	filename_new := strings.Split(filename, ".md")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var end_text_page_article string = ""
	var title_overview_articles string = ""
	var description_overview_articles string = ""
	//var image_overview_articles string = ""
	for scanner.Scan() {
		content := scanner.Text()
		if strings.Contains(content, "###") {
			part := strings.Split(content, `###`)
			end_text_page_article += "<h3>" + part[1] + "</h3>"

		} else if strings.Contains(content, "##") {
			part := strings.Split(content, `##`)
			end_text_page_article += "<h2>" + part[1] + "</h2>"
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			end_text_page_article += "<h1>" + part[1] + "</h1>"
			title_overview_articles = "<h3>" + part[1] + "</h3>"
		} else if strings.Contains(content, "HEADIMAGE:") {
			part := strings.Split(content, `HEADIMAGE:`)
			end_text_page_article += "<img src='/pictures/" + part[1] + "'>"
			//	image_overview_articles = "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "DESCRIPTION:") {
			part := strings.Split(content, `DESCRIPTION:`)
			description_overview_articles = "<p>" + part[1] + "</p>"
		} else if strings.Contains(content, "[img]:") {
			part := strings.Split(content, `[img]:`)
			end_text_page_article += "<img src='/pictures/" + part[1] + "'>"
		} else if strings.Contains(content, "DATE:") {
			part := strings.Split(content, "DATE:")
			end_text_page_article += "<i>" + part[1] + "</i><br>"
		} else if strings.Contains(content, "TAGS:") {
			part := strings.Split(content, `TAGS:`)
			part_2 := strings.Split(part[1], ",")
			part_tags := "<articles-tags>"
			for i := 0; i < len(part_2); i++ {
				part_tags += "<a href='/tags/" + part_2[i] + "'>" + part_2[i] + "</a>"
				//	tag_collection += "#" + part_2[i] + ":" + filename_new[0] + ":" + title_overview_articles + ":" + description_overview_articles + ":" + image_overview_articles
			}
			part_tags += "</articles-tags>"

		} else if content == "" {
			end_text_page_article += "<br>"
		} else {
			end_text_page_article += "<p>" + content + "</p>"
		}

	}

	if action == "MainPageArticles" {
		MainPageArticles(direc, filename_new[0], title_overview_articles, description_overview_articles)
	} else if action == "MainPageProjects" {
		MainPageProjects(direc, filename_new[0], title_overview_articles, description_overview_articles)
	} else if action == "ArticleSinglePages" {
		page_Article(filename_new[0], end_text_page_article)
	} else if action == "ProjectsSinglePages" {
		page_Project(filename_new[0], end_text_page_article)
	} else if action == "ArticleMainPage" {
		page_Collection_Overview_Articles(filename_new[0], title_overview_articles, description_overview_articles)
	} else if action == "ProjectMainPage" {
		page_Collection_Overview_Projects(filename_new[0], title_overview_articles, description_overview_articles)
	}

	//page_Collection_Overview_Articles(filename_new[0], title_overview_articles, image_overview_articles, description_overview_articles, tags_page_article)*/

}

func MainPageArticles(derec string, filename string, title string, description string) {
	main_page_articles_element += "<a href='/articles/" + filename + "'+>" + title + description + "</a>"
}
func MainPageProjects(derec string, filename string, title string, description string) {
	main_page_projects_element += "<a href='/projects/" + filename + "'+>" + title + description + "</a>"
}

func page_Collection_Overview_Articles(filename string, title string, description string) {
	data_articles_index += "<a class='special-article' href='/articles/" + filename + "/index.html'>" + title + description + "</a>"
}

func page_Collection_Overview_Projects(filename string, title string, description string) {
	data_projects_index += "<a class='special-project' href='/projects/" + filename + "/index.html'>" + title + description + "</a>"
}
func page_Project(filename string, text string) {
	err := os.Mkdir("./httpdocs/projects/"+filename, 0755)
	err = os.WriteFile("./httpdocs/projects/"+filename+"/index.html", []byte(article_project_page_top+"<article-details>"+text+"</article-details>"+article_project_page_bottom), 0755)
	if err != nil {
		panic(err)
	}
}
func page_Article(filename string, text string) {
	err := os.Mkdir("./httpdocs/articles/"+filename, 0755)
	err = os.WriteFile("./httpdocs/articles/"+filename+"/index.html", []byte(article_project_page_top+"<article-details>"+text+"</article-details>"+article_project_page_bottom), 0755)
	if err != nil {
		panic(err)
	}
}

func readFileOthers(filename string) {
	file, err := os.Open("./other/" + filename)
	if err != nil {
		log.Fatalf("not opening")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var new_text string = ""
	for scanner.Scan() {
		content := scanner.Text()
		var end_type string = ""

		if strings.Contains(content, "###") {
			part := strings.Split(content, `###`)
			end_type = "<h3>" + part[1] + "</h3>"
		} else if strings.Contains(content, "##") {
			part := strings.Split(content, `##`)
			end_type = "<h2>" + part[1] + "</h2>"
		} else if strings.Contains(content, "#") {
			part := strings.Split(content, `#`)
			end_type = "<h1>" + part[1] + "</h1>"
		} else if content == "" {
			end_type = "<br>"
		} else {
			end_type = "<p>" + content + "</p>"
		}
		new_text += end_type

	}
	filename_new := strings.Split(filename, ".md")
	err = os.Mkdir("./httpdocs/"+filename_new[0], 0755)
	err = os.WriteFile("./httpdocs/"+filename_new[0]+"/index.html", []byte(article_project_page_top+new_text+article_project_page_bottom), 0755)
	if err != nil {
		panic(err)
	}
}
