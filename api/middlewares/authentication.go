package middlewares

import (
	"cloudstorageapi.com/api/helper"
	"cloudstorageapi.com/api/response"
	"cloudstorageapi.com/configs"
	"cloudstorageapi.com/models/spaceModel"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func AuthenticateUser(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//Authorization header [Space-ID, Access-Token"]
		accessToken := r.Header.Get("Access-Token")

		spaceIdStr := r.Header.Get("Space-ID")
		//Atempting to convert space id to int
		spaceId, err := strconv.Atoi(spaceIdStr)
		if err != nil {
			helper.NewJsonResponse(w).
				SetStatus(http.StatusBadRequest).
				Encode(response.NewErrorResponse(
					993, "Authorization failed due to invalid spaceid"))
			return
		}

		//Atempt to match security info from redis cache for better speed
		rc, rderr := configs.GetRedisConnection()
		if rderr == nil {
			defer rc.Close() //close on return
			rcToken, err := rc.Do("GET", "CSSPACEID"+spaceIdStr)
			if err == nil {
				rcToken, ok := rcToken.(byte)
				if ok && accessToken == string(rcToken) {
					h(w, r, ps)
					return
				} else if ok {
					goto fail
				}
			} else {
				log.Fatal(rderr)
			}
		}

		//Redis may not contain the key so Attempt to get from model (postgreSQL)
		if space, merr := spaceModel.FindSpaceById(spaceId); space == nil {
			helper.NewJsonResponse(w).
				SetStatus(http.StatusForbidden).
				Encode(response.NewFailResponse("Space id doesn't exists"))
			return
		} else if merr == nil {
			//checking access token
			if space.AccessToken == accessToken {
				if rderr == nil {
					//update redis cache for better speed
					defer rc.Do("SET",
						"CSSPACEID"+spaceIdStr, space.AccessToken)
				}
				h(w, r, ps)
				return
			}
		} else {
			helper.NewJsonResponse(w).
				SetStatus(http.StatusInternalServerError).
				Encode(response.NewFailResponse("Some problem occured"))
			return
		}

	fail:
		//Authentication failed
		helper.NewJsonResponse(w).
			SetStatus(http.StatusUnauthorized).
			Encode(response.NewFailResponse("Authorization failed"))
	}
}
