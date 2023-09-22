package tree

import (
	"comments/pkg/models"
)

// ArrayToTree принимает массив комментариев и возвращает массив
// комментариев в виде дерева
func ArrayToTree(comments []models.Comment) []models.CommentTree {
	var res []models.CommentTree

	// изначальный массив из комментариев без ParentID
	for _, commnt := range comments {
		if !commnt.ParentID.Valid {
			res = append(res, models.CommentTree{
				ID:               commnt.ID,
				NewsID:           commnt.NewsID,
				CreatedAt:        commnt.CreatedAt,
				Content:          commnt.Content,
				ThreadedComments: nil,
			})
		}
	}

	// рекурсивно добавлются вложенные комментарии
	for i, comment := range res {
		insertThread(&comment, &comments)
		res[i].ThreadedComments = comment.ThreadedComments
	}

	return res
}

func insertThread(tree *models.CommentTree, comments *[]models.Comment) {
	for _, comment := range *comments {
		if comment.ParentID.UUID == tree.ID && comment.ParentID.Valid {
			tree.ThreadedComments = append(tree.ThreadedComments, models.CommentTree{
				ID:               comment.ID,
				NewsID:           comment.NewsID,
				CreatedAt:        comment.CreatedAt,
				Content:          comment.Content,
				ThreadedComments: nil,
			})
		}
	}
	for i, threadComment := range tree.ThreadedComments {
		insertThread(&threadComment, comments)
		tree.ThreadedComments[i] = threadComment
	}
}
