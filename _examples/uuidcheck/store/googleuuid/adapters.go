// Code generated by sqlc-connect (https://github.com/walterwanderley/sqlc-connect). DO NOT EDIT.

package googleuuid

import (
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "uuidcheck/api/googleuuid/v1"
)

func toProduct(in Product) *pb.Product {

	out := new(pb.Product)
	out.Id = in.ID
	if in.Category.Valid {
		out.Category = wrapperspb.Int32(in.Category.Int32)
	}
	if in.Name.Valid {
		out.Name = wrapperspb.String(in.Name.String)
	}
	return out
}
