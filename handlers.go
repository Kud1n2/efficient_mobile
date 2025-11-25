package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSubscriptions godoc
// @Summary Get all subscriptions
// @Description Retrieve all subscriptions from database
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} WithDate
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /effective_mobile [get]
func getAll(context *gin.Context) {
	rows, err := db.Query(
		"SELECT * FROM subscriptions",
	)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't add to DB."})
		log.Println("Couldn't add to DB.")
		return
	}
	defer rows.Close()
	var withDates []WithDate
	for rows.Next() {
		var name WithDate
		var id int
		err := rows.Scan(&id, &name.Service_name, &name.Price, &name.User_id, &name.Start_date, &name.Finish_date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			log.Println("Bad GET-request.")
			return
		}
		withDates = append(withDates, name)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		log.Printf("DB Error.")
		return
	}
	context.IndentedJSON(http.StatusOK, withDates)
	log.Println("Outputed all elements.")
}

// CreateSubscription godoc
// @Summary Create new subscription
// @Description Add a new subscription to database
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param subscription body WithDate true "Subscription data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /effective_mobile [post]
func post(context *gin.Context) {
	var name WithDate
	err := context.BindJSON(&name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		log.Println("Bad POST-request.")
		return
	}
	_, err = db.Exec(
		"INSERT INTO subscriptions (service_name, price, user_id, start_date, finish_date) VALUES ($1, $2,$3,$4,$5)",
		name.Service_name, name.Price, name.User_id, name.Start_date, name.Finish_date,
	)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't add to DB."})
		log.Println("Couldn't add to DB.")
		return
	}
	log.Printf("Created new element: %s\n", name.Service_name)
}

// GetSubscriptionByName godoc
// @Summary Get subscription by service name
// @Description Retrieve subscription by service name using path parameter
// @Tags Subscriptions
// @Produce json
// @Param service_name path string true "Service Name"
// @Success 200 {array} WithDate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /effective_mobile/{service_name} [get]
func getByName(context *gin.Context) {
	var name FindByName
	err := context.BindJSON(&name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request 1"})
		log.Println("Bad GET-request.")
		return
	}
	rows, err := db.Query("SELECT * FROM subscriptions WHERE service_name = $1", name.Service_name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request 2"})
		log.Println("Bad GET-request.")
	}
	var withDates []WithDate
	for rows.Next() {
		var name WithDate
		var id int
		err := rows.Scan(&id, &name.Service_name, &name.Price, &name.User_id, &name.Start_date, &name.Finish_date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request 3"})
			log.Println("Bad GET-request.")
			return
		}
		withDates = append(withDates, name)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		log.Printf("DB Error.")
		return
	}
	defer rows.Close()
	if len(withDates) == 0 {
		log.Printf("No such element.")
		context.JSON(http.StatusBadRequest, gin.H{"message": "No such element"})
		return
	}
	context.IndentedJSON(http.StatusOK, withDates)
	log.Printf("Found element %s. \n", name.Service_name)
}

// UpdateSubscription godoc
// @Summary Update subscription
// @Description Update existing subscription information
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param subscription body UpdateByName true "Updated subscription data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /effective_mobile/id [put]
func updateByServiceName(context *gin.Context) {
	var name UpdateByName
	err := context.BindJSON(&name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request 1"})
		log.Println("Bad PUT-request.")
		return
	}
	_, err = db.Exec("UPDATE subscriptions SET service_name = $1, price = $2, user_id = $3, start_date = $4, finish_date = $5 WHERE service_name = $6", name.Service_name, name.Price, name.User_id, name.Start_date, name.Finish_date, name.Old_service_name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't update element in DB."})
		log.Println("Couldn't update element in DB.")
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated"})
	log.Printf("Element %s updated.\n", name.Old_service_name)
}

// DeleteSubscription godoc
// @Summary Delete subscription by service name
// @Description Remove subscription from database by service name
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param service_name body FindByName true "Service name to delete"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /effective_mobile/id [delete]
func deleteByServiceName(context *gin.Context) {
	var name FindByName
	err := context.BindJSON(&name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		log.Println("Bad DELETE-request.")
		return
	}
	_, err = db.Exec("DELETE FROM subscriptions WHERE service_name = $1", name.Service_name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't update element in DB."})
		log.Println("Couldn't update element in DB.")
		return
	}
}

// CalculateTotalCost godoc
// @Summary Calculate total subscription costs
// @Description Calculate sum of subscription prices with optional filters
// @Tags Analytics
// @Accept json
// @Produce json
// @Param filters body GetSum true "Filter criteria"
// @Success 200 {object} SumResponse
// @Failure 400 {object} ErrorResponse
// @Router /getSum [get]
func getSum(context *gin.Context) {
	var filter GetSum
	rows, err := db.Query(
		"SELECT * FROM subscriptions",
	)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't add to DB."})
		log.Println("Couldn't add to DB.")
		return
	}
	defer rows.Close()
	var withDates []WithDate
	for rows.Next() {
		var name WithDate
		var id int
		err := rows.Scan(&id, &name.Service_name, &name.Price, &name.User_id, &name.Start_date, &name.Finish_date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			log.Println("Bad GET-request.")
			return
		}
		withDates = append(withDates, name)
	}
	if err := rows.Err(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		log.Printf("DB Error.")
		return
	}
	sum := 0
	err = context.BindJSON(&filter)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		log.Println("Bad GET-request.")
		return
	}
	DateStart, err := parseTime(filter.Start_date)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		log.Println("Bad start date parsing.")
		return
	}
	DateFinish, err := parseTime(filter.Finish_date)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		log.Println("Bad finish date parsing.")
		return
	}
	for _, x := range withDates {
		xDateStart, err := parseTime(x.Start_date)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			log.Println("Bad X start date parsing.")
			return
		}
		if x.Finish_date != "" {
			xDateFinish, err := parseTime(x.Finish_date)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
				log.Println("Bad X finish date parsing.")
				return
			}
			if filter.Service_name == "" {
				if filter.User_id == "" {
					if (DateStart.Before(xDateFinish) || DateStart.Equal(xDateFinish)) && (DateFinish.After(xDateStart) || DateFinish.Equal(xDateStart)) {
						sum += x.Price
					}
				} else {
					if (DateStart.Before(xDateFinish) || DateStart.Equal(xDateFinish)) && (DateFinish.After(xDateStart) || DateFinish.Equal(xDateStart)) && filter.User_id == x.User_id {
						sum += x.Price
					}
				}
			} else {
				if filter.User_id == "" {
					if (DateStart.Before(xDateFinish) || DateStart.Equal(xDateFinish)) && (DateFinish.After(xDateStart) || DateFinish.Equal(xDateStart)) && filter.Service_name == x.Service_name {
						sum += x.Price
					}
				} else {
					if (DateStart.Before(xDateFinish) || DateStart.Equal(xDateFinish)) && (DateFinish.After(xDateStart) || DateFinish.Equal(xDateStart)) && filter.User_id == x.User_id && filter.Service_name == x.Service_name {
						sum += x.Price
					}
				}
			}

		} else {
			if filter.Service_name == "" {
				if filter.User_id == "" {
					if (DateStart.Before(xDateStart) || DateStart.Equal(xDateStart)) && (xDateStart.Before(DateFinish) || DateFinish.Equal(xDateStart)) {
						sum += x.Price
					}
				} else {
					if (DateStart.Before(xDateStart) || DateStart.Equal(xDateStart)) && (xDateStart.Before(DateFinish) || DateFinish.Equal(xDateStart)) && filter.User_id == x.User_id {
						sum += x.Price
					}
				}
			} else {
				if filter.User_id == "" {
					if (DateStart.Before(xDateStart) || DateStart.Equal(xDateStart)) && (xDateStart.Before(DateFinish) || DateFinish.Equal(xDateStart)) && filter.Service_name == x.Service_name {
						sum += x.Price
					}
				} else {
					if (DateStart.Before(xDateStart) || DateStart.Equal(xDateStart)) && (xDateStart.Before(DateFinish) || DateFinish.Equal(xDateStart)) && filter.User_id == x.User_id && filter.Service_name == x.Service_name {
						sum += x.Price
					}
				}
			}
		}
	}
	log.Printf("Found sum is %d", sum)
	context.JSON(http.StatusOK, gin.H{"Sum": sum})
}
