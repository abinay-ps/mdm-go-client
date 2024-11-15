/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListReportsArticlesAbstractResponse struct {
	ArticleReceivedToday []DomainArticleTransaction `json:"article_received_today,omitempty"`
	ArticlesDeliveredByPostman []DomainArticleTransaction `json:"articles_delivered_by_postman,omitempty"`
	ArticlesDeliveredThroughWindow []DomainArticleTransaction `json:"articles_delivered_through_window,omitempty"`
	ArticlesDepositedCurrentDate []DomainArticleTransaction `json:"articles_deposited_current_date,omitempty"`
	ArticlesIssuedToBo []DomainArticleTransaction `json:"articles_issued_to_bo,omitempty"`
	ArticlesIssuedToPostman []DomainArticleTransaction `json:"articles_issued_to_postman,omitempty"`
	ArticlesRedirected []DomainArticleTransaction `json:"articles_redirected,omitempty"`
	ArticlesReturnedToSender []DomainArticleTransaction `json:"articles_returned_to_sender,omitempty"`
	DepositedArticles []DomainArticleTransaction `json:"deposited_articles,omitempty"`
	UnprocessedArticles []DomainUprocessArticle `json:"unprocessed_articles,omitempty"`
}
