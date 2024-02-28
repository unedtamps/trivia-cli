package repository

import "context"

func (q *Queries) TrviaCreateWithTx(
	ctx context.Context,
	arg CreateTriviaParams,
	opts []string,
) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	user_id := ctx.Value("id").(int64)
	id, err := q.CreateTrivia(ctx, arg)
	if err != nil {
		return err
	}
	for _, v := range opts {
		err = q.CreateTriviaChoice(ctx, CreateTriviaChoiceParams{
			TriviaID: id,
			Choice:   v,
		})
		if err != nil {
			return err
		}
	}
	err = q.CreateTriviaDetail(ctx, CreateTriviaDetailParams{
		UserID:         user_id,
		TriviaID:       id,
		TriviaStatusID: REVIEW,
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

type Triva struct {
	Title      string
	Question   string
	Aswere     string
	Dificulity int64
}

type Choice struct {
	Choice string
	Id     int64
}

func (q *Queries) UpdateTrivaWithTx(
	ctx context.Context,
	id int64,
	trivia Triva,
	choice []Choice,
) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := q.UpdateTrivia(ctx, UpdateTriviaParams{
		ID:           id,
		Title:        trivia.Title,
		Answer:       trivia.Aswere,
		Question:     trivia.Question,
		DificulityID: trivia.Dificulity,
	}); err != nil {
		return err
	}
	for _, v := range choice {
		if err := q.UpdateTriviaChoice(ctx, UpdateTriviaChoiceParams{
			ID:     v.Id,
			Choice: v.Choice,
		}); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (q *Queries) DeleteTrviaWithTx(ctx context.Context, id int64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// delete choice
	err = q.DeleteTriviaChoice(ctx, id)
	if err != nil {
		return err
	}
	err = q.DeleteTriviaDetail(ctx, id)
	if err != nil {
		return err
	}
	err = q.DeleteTrivia(ctx, id)
	if err != nil {
		return err
	}
	return tx.Commit()
}
