package tweet

import (
	"github.com/HotPotatoC/twitter-clone/internal/aws"
	"github.com/HotPotatoC/twitter-clone/internal/cache"
	"github.com/HotPotatoC/twitter-clone/internal/database"
	"github.com/HotPotatoC/twitter-clone/module/tweet/action"
	"github.com/HotPotatoC/twitter-clone/module/tweet/service"
	"github.com/HotPotatoC/twitter-clone/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, s3 *aws.S3Bucket, cache cache.Cache) {
	authMiddleware := middleware.NewAuthMiddleware()
	r.Get("/", buildListTweetHandler(db))
	r.Get("/feed", authMiddleware.Execute(), buildListTweetFeedHandler(db))
	r.Get("/search", authMiddleware.Execute(), buildSearchTweetHandler(db))
	r.Get("/:tweetID", authMiddleware.Execute(), buildGetTweetHandler(db))
	r.Post("/", authMiddleware.Execute(), buildCreateTweetHandler(db, s3))
	r.Get("/:tweetID/replies", authMiddleware.Execute(), buildListTweetRepliesHandler(db))
	r.Post("/:tweetID/reply", authMiddleware.Execute(), buildCreateReplyHandler(db))
	r.Post("/:tweetID/favorite", authMiddleware.Execute(), buildFavoriteTweetHandler(db))
}

func buildListTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewListTweetService(db)
		action := action.NewListTweetAction(service)

		return action.Execute(c)
	}
}

func buildListTweetFeedHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewListTweetFeedService(db)
		action := action.NewListTweetFeedAction(service)

		return action.Execute(c)
	}
}

func buildGetTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewGetTweetService(db)
		action := action.NewGetTweetAction(service)

		return action.Execute(c)
	}
}

func buildSearchTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewSearchTweetService(db)
		action := action.NewSearchTweetAction(service)

		return action.Execute(c)
	}
}

func buildCreateTweetHandler(db database.Database, s3 *aws.S3Bucket) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewCreateTweetService(db, s3)
		action := action.NewCreateTweetAction(service)

		return action.Execute(c)
	}
}

func buildListTweetRepliesHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewListTweetRepliesService(db)
		action := action.NewListTweetRepliesAction(service)

		return action.Execute(c)
	}
}

func buildCreateReplyHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewCreateReplyService(db)
		action := action.NewCreateReplyAction(service)

		return action.Execute(c)
	}
}

func buildFavoriteTweetHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewFavoriteTweetService(db)
		action := action.NewFavoriteTweetAction(service)

		return action.Execute(c)
	}
}