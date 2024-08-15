package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImages(paths []string) []*Job {
	var jobs []*Job

	// For each input path create job struct and add it to the list
	for _, p := range paths {
		job := &Job{
			InputPath: p,
			OutPath:   strings.Replace(p, "images/", "images/output/", 1),
		}
		job.Image = ReadImage(p)
		jobs = append(jobs, job)
	}

	return jobs
}

func resizeJobs(jobs []*Job) <-chan *Job {
	// For each input job, change the image in the job struct
	output := make(chan *Job, len(jobs))

	for _, j := range jobs {
		// j.Image = Resize(j.Image)
		go func(job *Job) {
			job.Image = Resize(job.Image)
			output <- job
		}(j)
	}

	return output
}

func saveImages(jobs []*Job) {
	for _, job := range jobs {
		WriteImage(job.OutPath, job.Image)
	}
}

func collectJobs(input <-chan *Job, count int) []*Job {
	resizedJobs := make([]*Job, 0, count)

	for i := 0; i < count; i++ {
		// read the job from the input channel
		job := <-input
		resizedJobs = append(resizedJobs, job)
	}

	return resizedJobs
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	imagePaths := []string{dir + "/images/image1.jpeg",
		dir + "/images/image2.jpeg",
		dir + "/images/image3.jpeg",
		dir + "/images/image4.jpeg",
	}

	jobs := loadImages(imagePaths)
	// Fan out this function to multiple goroutines
	out := resizeJobs(jobs)
	// Collect / Fan in result
	outJobs := collectJobs(out, len(jobs))

	saveImages(outJobs)
}
