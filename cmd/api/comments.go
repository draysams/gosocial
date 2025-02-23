package main

import (
	"net/http"

	"github.com/draysams/gosocial/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=500"`
}

func (app *application) createCommentOnPostHandler(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(postContextKey).(*store.Post)

	var payload CreateCommentPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	userId := int64(2)

	comment := &store.Comment{
		Content: payload.Content,
		PostID:  post.ID,
		UserID:  userId,
	}

	ctx := r.Context()
	if err := app.store.Comments.Create(ctx, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.writeJSONResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getCommentsOnPostHandler(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value(postContextKey).(*store.Post)

	ctx := r.Context()
	comments, err := app.store.Comments.GetCommentsByPostID(ctx, post.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.writeJSONResponse(w, http.StatusOK, comments); err != nil {
		app.internalServerError(w, r, err)
	}
}
