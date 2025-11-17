package main

const (
	ExamPageUrl       = "https://app.meijertheoriecursus.nl/app/examens"
	ExamenButtonLink  = ".//div[@class='grid gap-2 sm:grid-cols-[auto_120px]']/a"
	ExamStartButton   = ".//button[text()='Starten']"
	ExamAsset         = ".//img[contains(@class, 'mx-auto object-contain')] | .//video[contains(@class, 'mx-auto')]"
	ExamQuestion      = ".//h5[@class='font-heading font-bold text-lg md:text-xl']"
	ExamAnswerOptions = ".//div[contains(@class,'grid gap-2')]//span[@class='relative top-[1px]'] | .//div[contains(@class,'grid gap-2')]//span[contains(@class, 'ml-2')]"
	ExamNextQuestion  = ".//div[span[text()='Volgende vraag']]/button"
)

const (
	VideoCoursePageUrl = "https://app.meijertheoriecursus.nl/app/videocursus"
	VideoCourseButton  = ".//div/a[span]"
)

func main() {

}
