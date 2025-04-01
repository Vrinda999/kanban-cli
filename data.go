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
			title:       "Buy Milk",
			description: "Strawberry Milk",
		},
		Task{
			status:      todo,
			title:       "Eat Sushi",
			description: "Miso Soup, and Rice",
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
			title:       "Go Proj",
			description: "CLI Kanban",
		},
	})

	// Init Done
	b.cols[done].list.Title = "Done"
	b.cols[done].list.SetItems([]list.Item{
		Task{
			status:      done,
			title:       "Portfolio",
			description: "React, JS, Tailwind",
		},
	})
}
