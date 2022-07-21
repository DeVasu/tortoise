package products

import (
	"fmt"
	"time"

	"github.com/DeVasu/tortoise/datasources/mysql/cashiers_db"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)



const (
	queryInsertProduct         = "INSERT INTO products(categoryId, name, image, price, stock, discountQty, discountType, discountResult, discountExpiredAt, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryListProducts 			= "SELECT * from products;"
	queryById = "SELECT * from products where id=?;"
	queryUpdateProduct = "UPDATE products SET categoryId=?, name=?, image=?, price=?, stock=? WHERE id = ?;"
	queryDeleteProduct = "DELETE FROM products WHERE id=?;"
	
)

func(p *Product) Delete() *rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryDeleteProduct)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to update user")
	}
	return nil
}

func(p *Product) Update() *rest_errors.RestErr {

	temp := &Product{
		Id : p.Id,
	}
	temp.GetById()

	if p.CategoryId != 0 {
		temp.CategoryId = p.CategoryId
	}
	if len(p.Name) != 0 {
		temp.Name = p.Name
	}
	if len(p.Image) != 0 {
		temp.Image = p.Image
	}
	if p.Price != 0 {
		temp.Price = p.Price
	}
	if p.Stock != 0 {
		temp.Stock = p.Stock
	}


	stmt, err := cashiers_db.Client.Prepare(queryUpdateProduct)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		temp.CategoryId,
		temp.Name,
		temp.Image,
		temp.Price,
		temp.Stock,
		temp.Id,
	)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to update user")
	}
	return nil
}

func(p *Product) GetById() *rest_errors.RestErr {
	stmt, err := cashiers_db.Client.Prepare(queryById)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer stmt.Close()

	result := stmt.QueryRow(p.Id)

	if err := result.Scan(&p.Id,
		&p.CategoryId,
		&p.Name,
		&p.Image,
		&p.Price,
		&p.Stock,
		&p.UpdatedAt,
		&p.CreatedAt,
		&p.Discount.Qty, 
		&p.Discount.Type,
		&p.Discount.Result,
		&p.Discount.ExpiredAt,
		); err != nil {
		return rest_errors.NewInternalServerError("error when tying to gett cashier")
	}

	return nil
}

func(product *Product) Create() *rest_errors.RestErr {

	product.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	product.UpdatedAt = product.CreatedAt

	stmt, err := cashiers_db.Client.Prepare(queryInsertProduct)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get category")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(product.CategoryId, product.Name, product.Image, product.Price, product.Stock, product.Discount.Qty, product.Discount.Type, product.Discount.Result, product.Discount.ExpiredAt, product.UpdatedAt, product.CreatedAt)
	if saveErr != nil {
		return rest_errors.NewInternalServerError("error when tying to save user")
	}

	productId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save user")
	}
	product.Id = productId

	return nil
}

func (p *Product) List() ([]Product, *rest_errors.RestErr) {
	stmt, err := cashiers_db.Client.Prepare(queryListProducts)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer stmt.Close()
	rows, err := stmt.Query() //update with limit and skip
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer rows.Close()

	results := make([]Product, 0)
	for rows.Next() {
		var temp Product
		if err := rows.Scan(&temp.Id,
			&temp.CategoryId,
			&temp.Name,
			&temp.Image,
			&temp.Price,
			&temp.Stock,
			&temp.UpdatedAt,
			&temp.CreatedAt,
			&temp.Discount.Qty, 
			&temp.Discount.Type,
			&temp.Discount.Result,
			&temp.Discount.ExpiredAt,
			); err != nil {
			return nil, rest_errors.NewInternalServerError("error when tying to gett cashier")
		}
		results = append(results, temp)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}