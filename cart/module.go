package cart

import (
	"encoding/gob"

	"flamingo.me/flamingo-commerce/cart/application"
	"flamingo.me/flamingo-commerce/cart/domain/cart"
	"flamingo.me/flamingo-commerce/cart/infrastructure"
	"flamingo.me/flamingo-commerce/cart/interfaces/controller"
	"flamingo.me/flamingo-commerce/cart/interfaces/templatefunctions"
	"flamingo.me/flamingo/framework/config"
	"flamingo.me/flamingo/framework/dingo"
	"flamingo.me/flamingo/framework/event"
	"flamingo.me/flamingo/framework/router"
	"flamingo.me/flamingo/framework/template"
)

type (
	// CartModule registers our profiler
	CartModule struct {
		RouterRegistry  *router.Registry `inject:""`
		UseInMemoryCart bool             `inject:"config:cart.useInMemoryCartServiceAdapters"`
		EnableCartCache bool             `inject:"config:cart.enableCartCache,optional"`
	}
)

// Configure module
func (m *CartModule) Configure(injector *dingo.Injector) {
	if m.UseInMemoryCart {
		injector.Bind((*infrastructure.CartStorage)(nil)).To(infrastructure.InMemoryCartStorage{}).AsEagerSingleton()
		injector.Bind((*cart.GuestCartService)(nil)).To(infrastructure.InMemoryGuestCartService{})
		injector.Bind((*cart.CustomerCartService)(nil)).To(infrastructure.InMemoryCustomerCartService{})
	}
	//Register Default EventPublisher
	injector.Bind((*application.EventPublisher)(nil)).To(application.DefaultEventPublisher{})

	m.RouterRegistry.Handle("cart.view", (*controller.CartViewController).ViewAction)
	m.RouterRegistry.Route("/cart", "cart.view")

	m.RouterRegistry.Handle("cart.add", (*controller.CartViewController).AddAndViewAction)
	m.RouterRegistry.Route("/cart/add/:marketplaceCode", `cart.add(marketplaceCode,variantMarketplaceCode?="",qty?="1",deliveryIntent?="")`)

	m.RouterRegistry.Handle("cart.updateQty", (*controller.CartViewController).UpdateQtyAndViewAction)
	m.RouterRegistry.Route("/cart/update/:id", `cart.updateQty(id,qty?="1")`)

	m.RouterRegistry.Handle("cart.deleteAllItems", (*controller.CartViewController).DeleteAllAndViewAction)
	m.RouterRegistry.Route("/cart/delete/all", `cart.deleteAllItems`)

	m.RouterRegistry.Handle("cart.deleteItem", (*controller.CartViewController).DeleteAndViewAction)
	m.RouterRegistry.Route("/cart/delete/:id", `cart.deleteItem(id)`)

	gob.Register(cart.Cart{})

	//DecoratedCart API:

	m.RouterRegistry.Handle("cart.api.get", (*controller.CartApiController).GetAction)
	m.RouterRegistry.Handle("cart.api.add", (*controller.CartApiController).AddAction)
	m.RouterRegistry.Handle("cart.api.applyVoucher", (*controller.CartApiController).ApplyVoucherAndGetAction)

	m.RouterRegistry.Route("/api/cart", "cart.api.get")
	m.RouterRegistry.Route("/api/cart/add/:marketplaceCode", `cart.api.add(marketplaceCode,variantMarketplaceCode?="",qty?="1",deliveryIntent?="")`)
	m.RouterRegistry.Route("/api/cart/applyvoucher/:couponCode", `cart.api.applyVoucher(couponCode)`)

	//Event
	injector.BindMulti((*event.Subscriber)(nil)).To(application.EventReceiver{})

	// TemplateFunction
	injector.BindMulti((*template.ContextFunction)(nil)).To(templatefunctions.GetCart{})
	injector.BindMulti((*template.ContextFunction)(nil)).To(templatefunctions.GetDecoratedCart{})

	injector.Bind((*cart.DeliveryInfoBuilder)(nil)).To(cart.DefaultDeliveryInfoBuilder{})

	if m.EnableCartCache {
		injector.Bind((*application.CartCache)(nil)).To(application.CartSessionCache{})
	}

}

// DefaultConfig enables inMemory cart service adapter
func (m *CartModule) DefaultConfig() config.Map {
	return config.Map{
		"cart": config.Map{
			"useInMemoryCartServiceAdapters": true,
		},
	}
}
