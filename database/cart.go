package database

import(
fmt
)

var (
	
	ErrCantfindProduct = errors.New("cant find product")
	ErrCanDecodeProducts = errors.New("cant find product")
	ErrUserIdIsNotVaild = errors.New("not a vaild user")
	ErrCantUpdateUser = errors.New("cannot add the prodct to the cart")
	ErrCantRemoveItemCart = errors.New("cannt remove this item from cart")
	ErrCantGetItem = errors.New("was unable to get item from item")
	ErrCantBuyCartItem = errors.New("cannot update purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productId})

	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productCart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}
	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$earch", Value: productCart}}}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}

func RemoveCartItem(){

}

func BuyItemFromCart(){s

}
func InstantBuyer(){

}