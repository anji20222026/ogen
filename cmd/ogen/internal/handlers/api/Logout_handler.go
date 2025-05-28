package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
	"github.com/go-chi/chi/v5"
	
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/ogenregex"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"fastadmin/internal/handler/api"
	"fastadmin/fastadmin"
)

func (s *ApiImpl) Logout(ctx context.Context) (api.LogoutRes, error) {
	return api.LogoutRes{}, nil
}
