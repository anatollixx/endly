package v1



import (
	"fmt"
"errors"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	vvc "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"	
)

/*autogenerated contract adapter*/

//ServiceAccountCreateRequest represents request
type ServiceAccountCreateRequest struct {
  service_ v1.ServiceAccountInterface
  Account *vvc.ServiceAccount
}

//ServiceAccountUpdateRequest represents request
type ServiceAccountUpdateRequest struct {
  service_ v1.ServiceAccountInterface
  Account *vvc.ServiceAccount
}

//ServiceAccountDeleteRequest represents request
type ServiceAccountDeleteRequest struct {
  service_ v1.ServiceAccountInterface
  Name string
  Options *metav1.DeleteOptions
}

//ServiceAccountDeleteCollectionRequest represents request
type ServiceAccountDeleteCollectionRequest struct {
  service_ v1.ServiceAccountInterface
  Options *metav1.DeleteOptions
  ListOptions metav1.ListOptions
}

//ServiceAccountGetRequest represents request
type ServiceAccountGetRequest struct {
  service_ v1.ServiceAccountInterface
  Name string
  Options metav1.GetOptions
}

//ServiceAccountListRequest represents request
type ServiceAccountListRequest struct {
  service_ v1.ServiceAccountInterface
  Opts metav1.ListOptions
}

//ServiceAccountWatchRequest represents request
type ServiceAccountWatchRequest struct {
  service_ v1.ServiceAccountInterface
  Opts metav1.ListOptions
}

//ServiceAccountPatchRequest represents request
type ServiceAccountPatchRequest struct {
  service_ v1.ServiceAccountInterface
  Name string
  Pt types.PatchType
  Data []byte
  Subresources []string
}


func init() {
	register(&ServiceAccountCreateRequest{})
	register(&ServiceAccountUpdateRequest{})
	register(&ServiceAccountDeleteRequest{})
	register(&ServiceAccountDeleteCollectionRequest{})
	register(&ServiceAccountGetRequest{})
	register(&ServiceAccountListRequest{})
	register(&ServiceAccountWatchRequest{})
	register(&ServiceAccountPatchRequest{})
}


func (r * ServiceAccountCreateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountCreateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Create(r.Account)
	return result, err	
}

func (r * ServiceAccountCreateRequest) GetId() string {
	return "v1.ServiceAccountInterface.Create";	
}

func (r * ServiceAccountUpdateRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountUpdateRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Update(r.Account)
	return result, err	
}

func (r * ServiceAccountUpdateRequest) GetId() string {
	return "v1.ServiceAccountInterface.Update";	
}

func (r * ServiceAccountDeleteRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountDeleteRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.Delete(r.Name,r.Options)
	return result, err	
}

func (r * ServiceAccountDeleteRequest) GetId() string {
	return "v1.ServiceAccountInterface.Delete";	
}

func (r * ServiceAccountDeleteCollectionRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountDeleteCollectionRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	err = r.service_.DeleteCollection(r.Options,r.ListOptions)
	return result, err	
}

func (r * ServiceAccountDeleteCollectionRequest) GetId() string {
	return "v1.ServiceAccountInterface.DeleteCollection";	
}

func (r * ServiceAccountGetRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountGetRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Get(r.Name,r.Options)
	return result, err	
}

func (r * ServiceAccountGetRequest) GetId() string {
	return "v1.ServiceAccountInterface.Get";	
}

func (r * ServiceAccountListRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountListRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.List(r.Opts)
	return result, err	
}

func (r * ServiceAccountListRequest) GetId() string {
	return "v1.ServiceAccountInterface.List";	
}

func (r * ServiceAccountWatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountWatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Watch(r.Opts)
	return result, err	
}

func (r * ServiceAccountWatchRequest) GetId() string {
	return "v1.ServiceAccountInterface.Watch";	
}

func (r * ServiceAccountPatchRequest) SetService(service interface{}) error {
	var ok bool
	if r.service_, ok = service.(v1.ServiceAccountInterface); !ok {
		return fmt.Errorf("invalid service type: %T, expected: v1.ServiceAccountInterface", service)
	}
	return nil
}

func (r * ServiceAccountPatchRequest) Call() (result interface{}, err error) {
	if r.service_ == nil {
		return nil, errors.New("service was empty")
	}
	result, err = r.service_.Patch(r.Name,r.Pt,r.Data,r.Subresources...)
	return result, err	
}

func (r * ServiceAccountPatchRequest) GetId() string {
	return "v1.ServiceAccountInterface.Patch";	
}