package main

import (
	"go_recipes/services"
	"log"
)

func main() {
	/*
		path, pathErr := filepath.Abs("./assets/")
		if pathErr != nil {
			log.Fatalln(pathErr)
		}

		fmt.Println(path) */

	hfService := services.NewHelloFreshPdf("assets")
	_, err := hfService.GetAllFilesToRead()

	if err != nil {
		log.Fatalln(err)
	}

	/* 	sql := repository.DbConnect(utils.GetEnvFile().Name)
	   	dishRepo := repository.DishRepository{Mysql: *sql}
	   	_, err := dishRepo.GetWithIngredients(1)
	   	if err != nil {
	   		log.Fatalln(err)
	   	} */

	/* 	sql := repository.DbConnect(utils.GetEnvFile().Name)
	   	ing := repository.IngRepository{Mysql: *sql}
	   	test, err := ing.GetOrCreate("test")
	   	if err != nil {
	   		log.Fatalln(err)
	   	}

	   	log.Printf("%+v\n", test) */
	/* 	err := pdf.CropFile(0, 0, -50, 0)
	   	if err != nil {
	   		log.Fatalln(err)
	   	} */

	/* 	pdf := pdftools.PdfToImport{
	   		FileName:  "test1",
	   		Extension: ".pdf",
	   		Path:      "./assets/",
	   	}

	   	useless := []string{
	   		"A ajouter vous-meme",
	   		"* Conserver au réfrigérateur",
	   		"Valeurs nutritionnelles",
	   		"Par portion Pour 100 g",
	   	}

	   	uselessStep := []string{
	   		"LETRI",
	   		"+ FACILE",
	   		"A vos fourchettes !",
	   	}

	   	readErr := pdf.ReadFile([]uint16{1}, []uint16{2, 3, 4, 6, 7, 8}, useless, uselessStep)
	   	if readErr != nil {
	   		log.Fatalln(readErr)
	   	} */

	/* 	splitErr := pdf.SplitFile(4, 2)
	   	if splitErr != nil {
	   		panic(splitErr)
	   	} */

	/* 	cmd := exec.Command("echo", "hello", "world")
	   	if errors.Is(cmd.Err, exec.ErrDot) {
	   		cmd.Err = nil
	   	} */

	/* 	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	} */
	/* 	stdout, outErr := cmd.Output()
	   	if outErr != nil {
	   		log.Fatalln(outErr)
	   	}

	   	fmt.Println(string(stdout))
	*/
	/*
		 	err := repository.Migrate()
			if err != nil {
				log.Fatalln(err)
			}
	*/
}
