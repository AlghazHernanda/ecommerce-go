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

	
}

func RemoveCartItem(){

}

func BuyItemFromCart(){s

}
func InstantBuyer(){

}