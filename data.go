package main

import "github.com/charmbracelet/bubbles/list"

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}

	// Init To Do
	b.cols[todo].list.Title = "To Do"
	b.cols[todo].list.SetItems([]list.Item{
		Task{
			status:      todo,
			title:       "Study for Exam",
			description: "Final Exams!!",
		},
		Task{
			status:      todo,
			title:       "Eat Dinner",
			description: "Pizza Today!",
		},
		Task{
			status:      todo,
			title:       "Fold Laundry",
			description: "or Wear Wrinkly Tees",
		},
	})

	// Init in Progress
	b.cols[inProgress].list.Title = "In Progress"
	b.cols[inProgress].list.SetItems([]list.Item{
		Task{
			status:      inProgress,
			title:       "Final Year Project",
			description: "Disk Forensics Tool",
		},
	})

	// Init Done
	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		Task{
			status:      done,
			title:       "Portfolio Website",
			description: "Made in React, JS, Tailwind",
		},
	})
}
